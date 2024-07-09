import React from "react";
import { Link } from "@inertiajs/react";

interface pageProps {
  teamName: string;
}

export default function Index({ teamName }: pageProps) {
  return (
    <div>
      <p>Time {teamName}!</p>
      <Link href="/user">Ver usuário</Link>
      <Link href="/providers">Operadoras</Link>
    </div>
  );
}
