export function runOnceOnIntersect(
  element: Element,
  callback: () => void,
  rootMargin: string = "0",
) {
  let animationStarted = false;
  const observer = new IntersectionObserver(
    function (entries) {
      if (animationStarted) return;
      if (entries[0].isIntersecting === false) return;

      animationStarted = true;
      observer.unobserve(element);
      callback();
    },
    {
      rootMargin: rootMargin,
    },
  );
  observer.observe(element);
}
