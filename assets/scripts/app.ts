import { scrollToAnchor } from "./scroll-to-achor";

for (const link of document.querySelectorAll("a")) {
  link.addEventListener("click", scrollToAnchor);
}

console.log(
  "Want to view the source for this site? Go to https://github.com/tylerlwsmith/go-bakersfieldtechnology"
);
