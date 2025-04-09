// components/HealthStatus.tsx

'use client'

import { useEffect, useState } from 'react'
import { HealthService } from '@services/HealthService'

export default function HealthStatus() {
  const [status, setStatus] = useState<string | null>(null)
  const [dbConnected, setDbConnected] = useState<boolean | null>(null)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const service = new HealthService()

    service.check()
      .then((res) => {
        setStatus(res.status)
        console.log(res)
        setDbConnected(true)
      })
      .catch((err) => {
        setError(err.message || 'Unknown error')
      })
  }, [])

  if (error) return <div>Error: {error}</div>
  if (status === null || dbConnected === null) return <div>Loading...</div>

  return (
    <div>
      <p>Status: {status}</p>
      <p>Database connected: {dbConnected ? 'Yes' : 'No'}</p>
    </div>
  )
}
