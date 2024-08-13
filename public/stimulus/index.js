import {
  Application,
  Controller,
} from "https://unpkg.com/@hotwired/stimulus/dist/stimulus.js";
window.Stimulus = Application.start();

Stimulus.register(
  "toggleVisibility",
  class ToggleVisibility extends Controller {
    static targets = ["content"];

    constructor() {
      super();

      this.listener = null;
    }

    connect() {
      this.contentTarget.addEventListener(
        "htmx:beforeOnLoad",
        this.hideContent
      );

      this.hideContent();
    }

    showContent() {
      this.contentTarget.classList.remove("!hidden");
    }

    hideContent() {
      this.contentTarget.classList.add("!hidden");
    }

    disconnect() {
      this.contentTarget.removeEventListener(
        "htmx:beforeOnLoad",
        this.hideContent
      );
    }
  }
);
