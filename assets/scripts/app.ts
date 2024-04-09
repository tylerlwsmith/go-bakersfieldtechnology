import { scrollToAnchor } from "./scroll-to-achor";
import { triggerFadeInGroup } from "./trigger-fade-in-group";

console.log(
  "Want to view the source for this site? Go to https://github.com/tylerlwsmith/go-bakersfieldtechnology"
);

for (const link of document.querySelectorAll("a")) {
  link.addEventListener("click", scrollToAnchor);
}

triggerFadeInGroup(
  document.querySelectorAll(
    // [data-component="hero"]
    '[data-component="fade-in"][data-visible="false"]'
  )
);
