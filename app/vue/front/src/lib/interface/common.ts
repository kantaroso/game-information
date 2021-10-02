const toUnderscoreCase = (str: string): string => {
  return str.split(/(?=[A-Z])/).join('_').toLowerCase()
}

export const toUnderscoreCaseObject = (obj: any): any => { // eslint-disable-line @typescript-eslint/no-explicit-any
  const result: any = {}// eslint-disable-line @typescript-eslint/no-explicit-any
  Object.keys(obj).forEach(key => {
    result[toUnderscoreCase(key)] = obj[key]
  })
  return result
}
