import { Link } from "@inertiajs/react";
import React from "react";

interface pageProps {
  userName: string;
}

export default function User({ userName }: pageProps) {
  return (
    <div>
      <p>Usuário {userName}!</p>
      <Link href="/">Ver Time</Link>
    </div>
  );
}
