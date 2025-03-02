import type { HttpClient, Resp } from '../request'
import request from '../request'
import type { LoginToken } from '~/types'

class UserApi {
  private request: HttpClient

  constructor(request: HttpClient) {
    this.request = request
  }

  public async Login(password: string): Promise<Resp<LoginToken>> {
    return this.request.post('/user/login', { password })
  }

  public async UpdatePassword(password: string): Promise<Resp<null>> {
    return this.request.put('/user/password', { password })
  }

  public async UpdateToken(): Promise<Resp<LoginToken>> {
    return this.request.put('/user/token')
  }

  public async Verify(): Promise<Resp<null>> {
    return this.request.get('/user/verify')
  }
}

export default new UserApi(request)
