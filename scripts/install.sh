#!/bin/sh
# This script installs MeltCD on Linux, MacOS and WSL.
# It detects the current operating system architecture and installs the appropriate version of meltcd.

# ref: https://ollama.ai/install.sh

#set -eu

# Reset
Color_Off=''

# Regular Colors
Red=''
Green=''
Yellow=''
Dim='' # White

if [[ -t 1 ]]; then
    # Reset
    Color_Off='\033[0m' # Text Reset

    # Regular Colors
    Red='\033[1;31m'   # Red
    Green='\033[1;32m' # Green
    Yellow='\033[1;33m' # Yellow
    Dim='\033[1;2m'    # White
fi


status() { echo -e "${Dim}INFO   ${Color_Off}: $*" >&2; }
success() { echo -e "${Green}SUCCESS${Color_Off}: $*"; }
error() { echo -e "${Red}ERROR${Color_Off}  : $*"; exit 1; }
warning() { echo -e "${Yellow}WARNING${Color_Off}: $*"; }

TEMP_DIR=$(mktemp -d)
cleanp() { rm -rf $TEMP_DIR; }
trap cleanup EXIT

available() { command -v $1 >/dev/null; }
require() {
    local MISSING=''
    for TOOL in $*; do
        if ! available $TOOL; then
            MISSING="$MISSING $TOOL"
        fi
    done

    echo $MISSING
}

OS="$(uname -s)"
if [[ "$OS" =~ "Linux" ]]; then 
    OS="Linux"
elif [[ "$OS" =~ "Darwin" ]]; then
    OS="Darwin"
else 
    error 'This script is intended to run on Linux, MacOS and WSL only.'
fi

ARCH=$(uname -m)
case "$ARCH" in
    x86_64) ARCH="x86_64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) error "Unsupported architecture: $ARCH" ;;
esac

SUDO=
if [ "$(id -u)" -ne 0 ]; then
    # Running as root, no need for sudo
    if ! available sudo; then
        error "This script requires superuser permissions. Please re-run as root."
    fi

    SUDO="sudo"
fi

NEEDS=$(require curl tar grep cut)
if [ -n "$NEEDS" ]; then
    status "ERROR: The following tools are required but missing:"
    for NEED in $NEEDS; do
        echo "  - $NEED"
    done
    exit 1
fi

LATEST_VERSION=$(curl --fail --show-error --location -s https://api.github.com/repos/meltred/meltcd/releases/latest | grep "tag_name" | cut -d '"' -f 4)

status "Downloading MeltCD v$LATEST_VERSION for $OS $ARCH"
curl --fail --show-error --location --progress-bar -o $TEMP_DIR/meltcd.tar.gz "https://github.com/meltred/meltcd/releases/download/$LATEST_VERSION/meltcd_${LATEST_VERSION}_${OS}_$ARCH.tar.gz"

tar zxf $TEMP_DIR/meltcd.tar.gz -C $TEMP_DIR

for BINDIR in /usr/local/bin /usr/bin /bin; do
    echo $PATH | grep -q $BINDIR && break || continue
done

status "Installing MeltCD to $BINDIR..."
$SUDO install -o0 -g0 -m755 -d $BINDIR
$SUDO install -o0 -g0 -m755 $TEMP_DIR/meltcd $BINDIR/meltcd

install_success() { 
    [[ "$OS" = "Linux" ]] && success 'The MeltCD Server is running at 0.0.0.0:11771.'
    success 'Install complete. Run "meltcd" from the command line.'
}
trap install_success EXIT

configure_systemd() {
    if ! id meltcd >/dev/null 2>&1; then
        status "Creating meltcd user..."
        $SUDO useradd -r -s /bin/false -m -d /usr/share/meltcd meltcd
    fi

    status "Adding current user to meltcd group..."
    $SUDO usermod -a -G meltcd $(whoami)

    status "Creating meltcd systemd service..."
    cat <<EOF | $SUDO tee /etc/systemd/system/meltcd.service >/dev/null
[Unit]
Description=MeltCD Service
After=network-online.target

[Service]
ExecStart=$BINDIR/meltcd serve
User=meltcd
Group=meltcd
Restart=always
RestartSec=3
Environment="PATH=$PATH"

[Install]
WantedBy=default.target
EOF
    SYSTEMCTL_RUNNING="$(systemctl is-system-running || true)"
    case $SYSTEMCTL_RUNNING in
        running|degraded)
            status "Enabling and starting meltcd service..."
            $SUDO systemctl daemon-reload
            $SUDO systemctl enable meltcd
            $SUDO systemctl restart meltcd
            ;;
    esac
}


if available systemctl && available tee; then
    [[ "$OS" = "Linux" ]] && configure_systemd
fi
