package spec

import (
	"os"
	"testing"
)

func TestNormalizeFilePath(t *testing.T) {
	testCase := []string{
		".env",
		"file.txt",
		"~/.service.yaml",
	}

	for _, file := range testCase {
		res, err := normalizeFilePath(file)
		if err != nil {
			t.Error(err.Error())
		}
		// when we normalize filePath it become like
		// file.txt    /home/user/directory/file.txt
		// ~/file.txt /home/user/file.txt
		//
		// so the length should not equal to expanded result
		if len(res) == len(file) {
			t.Log("file does not example", file)
			t.Fail()
		}
	}
}

func TestGetEnvVars(t *testing.T) {
	tempFile, err := os.CreateTemp(os.TempDir(), "test_file-*")
	if err != nil {
		t.Error(err.Error())
	}
	defer tempFile.Close()

	tempFile.WriteString(`
	ENV_1="1"
	ENV_2="2"
	ENV_3=3
	ENV_4=4
	# Comment
	`)

	path := tempFile.Name()

	res, err := getEnvVars(path)
	if err != nil {
		t.Error(err.Error())
	}

	if res["ENV_1"] != "1" ||
		res["ENV_2"] != "2" ||
		res["ENV_3"] != "3" ||
		res["ENV_4"] != "4" {
		t.Error("failed to convert env file into map[string]string", res)
	}
}
