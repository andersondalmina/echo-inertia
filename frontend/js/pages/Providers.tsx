import React from "react";
import { Link } from "@inertiajs/react";

interface pageProps {
  providers: any[];
}

export default function Providers({ providers }: pageProps) {
  return (
    <div>
      {providers.map((provider) => (
        <p>Time {provider.name}!</p>
      ))}

      <Link href="/">Home</Link>
    </div>
  );
}
