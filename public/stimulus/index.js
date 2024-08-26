import {
  Application,
  Controller,
} from "https://unpkg.com/@hotwired/stimulus/dist/stimulus.js";
window.Stimulus = Application.start();

class ToggleVisibility extends Controller {
  static targets = ["content"];

  connect() {
    this.contentTarget.addEventListener("htmx:beforeOnLoad", () => {
      this.hideContent();
    });

    this.hideContent();
  }

  showContent() {
    this.contentTarget.classList.remove("!hidden");
  }

  hideContent() {
    this.contentTarget.classList.add("!hidden");
  }
}

Stimulus.register("toggleVisibility", ToggleVisibility);
