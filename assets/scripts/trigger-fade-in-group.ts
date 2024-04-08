export function triggerFadeInGroup(elements: NodeListOf<Element>) {
  let itemIndex = 0;

  setTimeout(function delayAnimationStart() {
    const visibilityInterval = setInterval(function showItems() {
      console.log("looping");
      if (itemIndex > elements.length - 1) {
        clearInterval(visibilityInterval);
        return;
      }

      elements[itemIndex].classList.remove("fade-in--start");
      elements[itemIndex].classList.add("fade-in--end");
      itemIndex++;
    }, 100);
  }, 100);
}
