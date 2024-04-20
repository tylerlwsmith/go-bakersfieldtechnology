import { scrollToAnchor } from "./scroll-to-achor";
import { triggerFadeInGroup } from "./trigger-fade-in-group";
import { runOnceOnIntersect } from "./run-once-on-intersect";

console.log(
  "Want to view the source for this site? Go to https://github.com/tylerlwsmith/go-bakersfieldtechnology",
);

const links = document.querySelectorAll("a");
for (const link of links) {
  link.addEventListener("click", scrollToAnchor);
}

const FADEIN_SELECTOR = '[data-component="fade-in"][data-visible="false"]';

const heroSection = document.querySelector('[data-component="hero"]');
if (heroSection !== null) {
  const heroFadeIns = heroSection.querySelectorAll(FADEIN_SELECTOR);

  window.addEventListener("load", () => triggerFadeInGroup(heroFadeIns));
}

const servicesSection = document.querySelector(
  '[data-component="services-list"]',
);
if (servicesSection !== null) {
  const servicesFadeIns = servicesSection.querySelectorAll(FADEIN_SELECTOR);
  runOnceOnIntersect(
    servicesSection,
    () => {
      triggerFadeInGroup(servicesFadeIns);
    },
    "-200px 0px",
  );
}
