const puppeteer = require('puppeteer');

// Gets the total request time based on the responseEnd - requestStart
// in secs
const getTotalRequestTime = (navigationEntries = []) => {
  if (navigationEntries.length === 0) {
    return null;
  }

  let firstRequestStart = null;
  let lastResponseEnd = null;

  navigationEntries.forEach((entry) => {
    if (entry.entryType === 'resource' || entry.entryType === 'navigation') {
      if (firstRequestStart === null || entry.requestStart < firstRequestStart) {
        firstRequestStart = entry.requestStart;
      }

      if (lastResponseEnd === null || entry.responseEnd > lastResponseEnd) {
        lastResponseEnd = entry.responseEnd;
      }
    }
  });

  return (lastResponseEnd - firstRequestStart) / 1000;
};

(async () => {
  const waitOptions = { waitUntil: 'networkidle0' };
  // for debugging
  const launchOptions = { headless: true };
  const browser = await puppeteer.launch(launchOptions);
  const page = await browser.newPage();

  // Set windows height en width
  await page.setViewport({ width: 1600, height: 900 });
  await page.goto('http://officelocal:3000/devlocal-auth/login', waitOptions);

  // Login by clicking button for existing user
  const loginBtnSelector = 'button[value="9bda91d2-7a0c-4de1-ae02-b8cf8b4b858b"]';
  await page.waitForSelector(loginBtnSelector);
  await Promise.all([page.click(loginBtnSelector), page.waitForNavigation(waitOptions)]);

  // hacky javascript timing
  const timeStart = Date.now();

  // go to a document viewer, orders
  await page.goto('http://officelocal:3000/moves/RBMDTK/orders', waitOptions);

  const pdfTitleSelector = 'div[data-testid="DocViewerContent"]';
  await page.waitForSelector(pdfTitleSelector);

  const timeEnd = Date.now();

  // Will return all http requests and navigation performance on last navigation
  const navigationEntries = JSON.parse(
    await page.evaluate(() => {
      return JSON.stringify(performance.getEntries());
    }),
  );

  // Getting some time discrepancies between the manual javascript timer and calculating the time from the entries
  // eslint-disable-next-line no-console
  console.log(getTotalRequestTime(navigationEntries), 'secs from performance entries');
  // eslint-disable-next-line no-console
  console.log((timeEnd - timeStart) / 1000, 'secs from javascript timer');
  await browser.close();
})();
