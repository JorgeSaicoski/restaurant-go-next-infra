"use client";

import HealthStatus from "@/components/HealthStatus";
import { useState } from "react";

export default function CounterButton() {
  const [count, setCount] = useState(0);

  return (
    <div>
      <button onClick={() => setCount(count + 1)}>Count: {count}</button>
      <HealthStatus></HealthStatus>
    </div>
  );
}
