export function initHeroFadeInAnimation() {
  // if (window.matchMedia("(prefers-reduced-motion: reduce)").matches) {
  //   return;
  // }
  const fadeInItems = document.querySelectorAll(".fade-in");
  let itemIndex = 0;

  console.log("Can you see me?");
  setTimeout(function delayAnimationStart() {
    const visibilityInterval = setInterval(function showItems() {
      console.log("looping");
      if (itemIndex > fadeInItems.length - 1) {
        clearInterval(visibilityInterval);
        return;
      }

      fadeInItems[itemIndex].classList.add("hidden");
      itemIndex++;
    }, 100);
  }, 100);
}
