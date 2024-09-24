export function formatDescription(text: string) {
  const lengthText = text.length

  if (lengthText > 10) {
    return `${text.substring(0, 10) + '...'}`
  } else {
    return text
  }
}
