export function scrollToAnchor(event: MouseEvent) {
  if (!(event.target instanceof HTMLAnchorElement)) return;

  const link = new URL(event.target.href);
  if (link.origin !== location.origin || link.pathname !== location.pathname) {
    return;
  }

  const targetId = link.hash;
  if (targetId === "") return;

  const section = document.querySelector(targetId);
  if (!(section instanceof HTMLElement)) return;

  event.preventDefault();
  const hasTabIndex = section.hasAttribute("tabindex");
  const hasOutline = section.style.outline ? true : false;

  if (!hasTabIndex) section.setAttribute("tabindex", "-1");
  if (!hasOutline) section.style.outline = "none";

  section.scrollIntoView({ block: "start" });
  section.focus({ preventScroll: true });

  section.addEventListener("blur", function blurListener() {
    section.removeEventListener("blur", blurListener);
    if (!hasTabIndex) section.removeAttribute("tabindex");
    if (!hasOutline) section.style.removeProperty("outline");
  });
}
