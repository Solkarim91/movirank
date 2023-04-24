import { test, expect } from '@playwright/test';

// EXAMPLE TESTS
test('has title', async ({ page }) => {
  await page.goto('https://playwright.dev/');

  // Expect a title "to contain" a substring.
  await expect(page).toHaveTitle(/Playwright/);
});

test('get started link', async ({ page }) => {
  await page.goto('https://playwright.dev/');

  // Click the get started link.
  await page.getByRole('link', { name: 'Get started' }).click();

  // Expects the URL to contain intro.
  await expect(page).toHaveURL(/.*intro/);
});

// ACTUAL TESTS
test.describe('Home page', () => {
  test('homepage has correct title', async ({ page }) => {
    await page.goto('http://localhost:3000/movies/');
  
    // Expect a title "to contain" a substring.
    await expect(page).toHaveTitle(/MoviRank/);
  });

  test('link to movie displays movie show page', async ({ page }) => {
    await page.goto('http://localhost:3000/movies/');
  
    // Locate the Link component with the text 'Titanic'
    const linkSelector = 'a:has-text("Titanic")';
    const linkElement = page.locator(linkSelector);
  
    // Click on the Link component
    await linkElement.click();
  
    // Wait for the selector that targets the text you want to check for
    const textSelector = 'div:has-text("James Cameron")';
    await page.waitForSelector(textSelector);
  
    // Retrieve the text content using page.textContent() method
    const textContent = await page.textContent(textSelector);
  
    // Make assertions on the text content
    expect(textContent).toContain('James Cameron');
  });
})
