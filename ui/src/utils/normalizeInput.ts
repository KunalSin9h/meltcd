import { toast } from "react-hot-toast";

export default function normalizeInput(
  givenText: string,
  allowed: string[],
  toastBottom?: boolean
): string {
  const len = givenText.length;
  let result = "";

  for (let i = 0; i < len; i++) {
    const code = givenText.charCodeAt(i);

    if (
      (code > 47 && code < 58) || // numeric (0-9)
      (code > 64 && code < 91) || // upper alpha (A-Z)
      (code > 96 && code < 123) || // lower alpha (a-z)
      allowed.includes(givenText[i])
    ) {
      result += givenText[i];
    } else {
      toast.error(`${givenText[i]} is not allowed in input here!`, {
        position: toastBottom ? "bottom-right" : "top-center",
      });
    }
  }

  return result;
}
