export function formatDate(isoString: string) {
  const date = new Date(isoString)

  const day = String(date.getDate()).padStart(2, '0')
  const month = String(date.getMonth() + 1).padStart(2, '0') // Months are 0-based
  const year = date.getFullYear()

  // Example format: '23/09/2024 dd/mm/yy'
  return `${day}/${month}/${year} `
}
