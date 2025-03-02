export interface LoginToken {
  token: string
}

export interface Device {
  id: number
  name: string
  mac: string
  delay: number
  topic: string
  password: string
  username: string
  ip: string
}
