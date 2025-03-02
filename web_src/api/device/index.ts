import type { HttpClient, Resp } from '../request'
import request from '../request'
import type { Device } from '~/types'

class DeviceApi {
  private request: HttpClient

  constructor(request: HttpClient) {
    this.request = request
  }

  public async AddDevice(device: Device): Promise<Resp<null>> {
    return this.request.post('/device', device)
  }

  public async GetDeviceList(): Promise<Resp<Device[]>> {
    return this.request.get('/device/list')
  }

  public async GetDeviceById(id: number): Promise<Resp<Device>> {
    return this.request.get(`/device/${id}`)
  }

  public async UpdateDevice(id: number, device: Device): Promise<Resp<null>> {
    return this.request.put(`/device/${id}`, device)
  }

  public async DeleteDevice(id: number): Promise<Resp<null>> {
    return this.request.delete(`/device/${id}`)
  }
}

export default new DeviceApi(request)
