import type { HttpClient, Resp } from '../request'
import request from '../request'

class WebApi {
  private request: HttpClient

  constructor(request: HttpClient) {
    this.request = request
  }

  public async Login(password: string): Promise<Resp<null>> {
    return this.request.post('/web/login')
  }
}

export default new WebApi(request)
