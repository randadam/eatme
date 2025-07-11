import { Page } from "@playwright/test";

export class SignupAccountPage {
  constructor(private page: Page) {}

  async goto() {
    await this.page.goto("/signup");
  }

  async fillEmail(email: string) {
    await this.page.getByRole("textbox", { name: /email/i }).fill(email);
  }

  async fillPassword(password: string) {
    await this.page.getByRole("textbox", { name: /password/i }).first().fill(password);
  }

  async fillConfirmPassword(password: string) {
    await this.page.getByRole("textbox", { name: /confirm password/i }).fill(password);
  }

  async submit() {
    await this.page.getByRole("button", { name: /next/i }).click();
  }
}
