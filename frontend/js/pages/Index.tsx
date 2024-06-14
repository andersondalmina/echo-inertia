import React from "react";

interface pageProps {
  teste: string
}

export default function Index({ teste }: pageProps) {
  return (
    <p>Hello {teste}, welcome to your first Inertia app!</p>
  );
}
