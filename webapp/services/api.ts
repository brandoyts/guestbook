class ApiService {
  private apiBaseEndpoint: string

  constructor() {
    const isDevelopment = process.env.NODE_ENV === "development"
    this.apiBaseEndpoint = isDevelopment
      ? (process.env.NEXT_PUBLIC_DEVELOPMENT_API_BASE_ENDPOINT as string)
      : (process.env.NEXT_PUBLIC_PRODUCTION_API_BASE_ENDPOINT as string)
  }

  async setItemToRedis(item: string) {
    const requestBody = JSON.stringify({ message: item })
    console.log(`${this.apiBaseEndpoint}/set`)
    try {
      const response = await fetch(`${this.apiBaseEndpoint}/set`, {
        headers: {
          "Content-Type": "application/json",
        },
        body: requestBody,
        method: "POST",
      })

      console.log(response)
    } catch (error) {
      console.log(error)
      throw error
    }
  }

  async getItemToRedis(itemKey: string) {
    try {
      const response = await fetch(`${this.apiBaseEndpoint}/get`, {
        headers: {
          "Content-Type": "application/json",
        },
      })

      console.log(response)
    } catch (error) {
      console.log(error)
      throw error
    }
  }
}

const Api = new ApiService()

export default Api
