import type { HttpClient, Resp } from '../request'
import request from '../request'

class ActionApi {
  private request: HttpClient

  constructor(request: HttpClient) {
    this.request = request
  }

  public async Ping(): Promise<Resp<null>> {
    return this.request.get('/action/ping?token=123456&name=pc')
  }
}

export default new ActionApi(request)
