export const fibonacci = (n: number): number => {
  if (n < 2) return n;

  return fibonacci(n - 1) + fibonacci(n - 2);
};

export const fibonacciMemorized = (n: number): number => {
  const cache = {};

  const f = (n: number) => {
    if (n < 2) return n;
    if (cache[n]) return cache[n];

    cache[n] = f(n - 1) + f(n - 2);
    return cache[n];
  };

  return f(n);
};
