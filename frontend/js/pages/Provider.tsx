import React from "react";
import { Link } from "@inertiajs/react";

interface pageProps {
  provider: any;
}

export default function Provider({ provider }: pageProps) {
  return (
    <div>
      <h1>{provider.name}</h1>

      <Link href="/">Home</Link>
    </div>
  );
}
