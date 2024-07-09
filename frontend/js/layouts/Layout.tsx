import React from "react";

interface layoutProps {
  children: React.ReactNode;
}

export default function Layout({ children }: layoutProps) {
  return (
    <main>
      <h1>Layout</h1>

      {children}
    </main>
  );
}
