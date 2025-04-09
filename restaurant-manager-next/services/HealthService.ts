const BASE_URL = process.env.NEXT_PUBLIC_GO_RESTAURANT_URL || "http://localhost:8082"

export class HealthService {
  private url: string = `${BASE_URL}/healthy`

  private headers: HeadersInit = {
    'Content-Type': 'application/json',
  }

  public async check(): Promise<{ status: string }> {
    try {
      const response = await fetch(this.url, {
        method: 'GET',
        headers: this.headers,
      })

      if (!response.ok) {
        throw new Error(`GET /healthy failed with status ${response.status}`)
      }

      return await response.json()
    } catch (error) {
      console.error("Health check failed:", error)
      throw error
    }
  }
}
