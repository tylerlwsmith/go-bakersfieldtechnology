export function triggerFadeInGroup(elements: NodeListOf<Element>) {
  let itemIndex = 0;

  setTimeout(function delayAnimationStart() {
    const visibilityInterval = setInterval(function showItems() {
      if (itemIndex > elements.length - 1) {
        clearInterval(visibilityInterval);
        return;
      }

      elements[itemIndex].setAttribute("data-visible", "true");
      itemIndex++;
    }, 100);
  }, 100);
}
