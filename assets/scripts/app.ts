import { scrollToAnchor } from "./scroll-to-achor";
import { initHeroFadeInAnimation } from "./hero";

console.log(
  "Want to view the source for this site? Go to https://github.com/tylerlwsmith/go-bakersfieldtechnology"
);

for (const link of document.querySelectorAll("a")) {
  link.addEventListener("click", scrollToAnchor);
}

initHeroFadeInAnimation();
