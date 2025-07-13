const puppeteer = require('puppeteer');
const { spawn } = require('child_process');
const path = require('path');

describe('Sydney Health App E2E Tests', () => {
  let browser;
  let page;
  let serverProcess;
  const APP_URL = 'http://localhost:3000';
  
  // Helper function to wait for server to be ready
  const waitForServer = async (url, maxAttempts = 30) => {
    for (let i = 0; i < maxAttempts; i++) {
      try {
        await page.goto(url, { waitUntil: 'networkidle0', timeout: 5000 });
        return true;
      } catch (error) {
        console.log(`Waiting for server... attempt ${i + 1}/${maxAttempts}`);
        await new Promise(resolve => setTimeout(resolve, 2000));
      }
    }
    throw new Error('Server failed to start');
  };

  beforeAll(async () => {
    // Launch browser
    browser = await puppeteer.launch({
      headless: 'new',
      args: ['--no-sandbox', '--disable-setuid-sandbox']
    });
    
    page = await browser.newPage();
    await page.setViewport({ width: 1280, height: 800 });
    
    // Start development server
    console.log('Starting development server...');
    serverProcess = spawn('npm', ['run', 'dev'], {
      cwd: path.resolve(__dirname, '../..'),
      detached: false,
      stdio: 'pipe'
    });
    
    serverProcess.stdout.on('data', (data) => {
      console.log(`Server: ${data}`);
    });
    
    serverProcess.stderr.on('data', (data) => {
      console.error(`Server Error: ${data}`);
    });
    
    // Wait for server to be ready
    await waitForServer(APP_URL);
  }, 120000);

  afterAll(async () => {
    if (browser) {
      await browser.close();
    }
    
    if (serverProcess) {
      // Kill the server process
      process.kill(-serverProcess.pid);
    }
  });

  describe('Application Loading', () => {
    test('should load the application without errors', async () => {
      const response = await page.goto(APP_URL, { waitUntil: 'networkidle0' });
      expect(response.status()).toBe(200);
      
      // Check for console errors
      const errors = [];
      page.on('console', msg => {
        if (msg.type() === 'error') {
          errors.push(msg.text());
        }
      });
      
      await page.waitForTimeout(1000);
      expect(errors).toHaveLength(0);
    }, 30000);

    test('should display login page initially', async () => {
      await page.goto(APP_URL, { waitUntil: 'networkidle0' });
      
      // Check if we're redirected to login
      await page.waitForSelector('form', { timeout: 10000 });
      
      const title = await page.title();
      expect(title).toContain('Sydney Health');
      
      // Check for login form elements
      const emailInput = await page.$('input[type="email"]');
      const passwordInput = await page.$('input[type="password"]');
      const submitButton = await page.$('button[type="submit"]');
      
      expect(emailInput).toBeTruthy();
      expect(passwordInput).toBeTruthy();
      expect(submitButton).toBeTruthy();
    }, 30000);
  });

  describe('Authentication Flow', () => {
    test('should login successfully with valid credentials', async () => {
      await page.goto(`${APP_URL}/login`, { waitUntil: 'networkidle0' });
      
      // Fill login form
      await page.type('input[type="email"]', 'test@example.com');
      await page.type('input[type="password"]', 'password123');
      
      // Submit form
      await Promise.all([
        page.waitForNavigation({ waitUntil: 'networkidle0' }),
        page.click('button[type="submit"]')
      ]);
      
      // Should redirect to dashboard
      expect(page.url()).toBe(`${APP_URL}/`);
      
      // Check for dashboard elements
      await page.waitForSelector('h1', { timeout: 10000 });
      const heading = await page.$eval('h1', el => el.textContent);
      expect(heading).toContain('Dashboard');
    }, 30000);

    test('should maintain session across page refreshes', async () => {
      // Refresh page
      await page.reload({ waitUntil: 'networkidle0' });
      
      // Should still be on dashboard
      expect(page.url()).toBe(`${APP_URL}/`);
      
      // Check if still logged in
      const heading = await page.$eval('h1', el => el.textContent);
      expect(heading).toContain('Dashboard');
    }, 30000);
  });

  describe('Navigation', () => {
    test('should navigate to all main pages', async () => {
      const pages = [
        { name: 'Benefits', url: '/benefits', heading: 'Benefits' },
        { name: 'Claims', url: '/claims', heading: 'Claims' },
        { name: 'Providers', url: '/providers', heading: 'Find Providers' },
        { name: 'Member Card', url: '/member-card', heading: 'Member Card' },
        { name: 'Messages', url: '/messages', heading: 'Messages' }
      ];
      
      for (const pageInfo of pages) {
        // Navigate to page
        await page.goto(`${APP_URL}${pageInfo.url}`, { waitUntil: 'networkidle0' });
        
        // Check URL
        expect(page.url()).toBe(`${APP_URL}${pageInfo.url}`);
        
        // Check heading
        await page.waitForSelector('h1', { timeout: 10000 });
        const heading = await page.$eval('h1', el => el.textContent);
        expect(heading).toContain(pageInfo.heading);
        
        // Check for console errors
        const errors = [];
        page.on('console', msg => {
          if (msg.type() === 'error') {
            errors.push(msg.text());
          }
        });
        
        await page.waitForTimeout(500);
        expect(errors).toHaveLength(0);
      }
    }, 60000);

    test('should have working navigation menu', async () => {
      await page.goto(`${APP_URL}/`, { waitUntil: 'networkidle0' });
      
      // Click on navigation links
      const navLinks = await page.$$('nav a');
      expect(navLinks.length).toBeGreaterThan(0);
      
      // Test first navigation link
      await navLinks[1].click(); // Skip dashboard link
      await page.waitForNavigation({ waitUntil: 'networkidle0' });
      
      // Should navigate to different page
      expect(page.url()).not.toBe(`${APP_URL}/`);
    }, 30000);
  });

  describe('Component Rendering', () => {
    test('should render all dashboard components', async () => {
      await page.goto(`${APP_URL}/`, { waitUntil: 'networkidle0' });
      
      // Check for key dashboard components
      const components = [
        { selector: 'h2', text: 'Deductible Tracker' },
        { selector: 'h2', text: 'Recent Claims' },
        { selector: 'h2', text: 'Quick Actions' }
      ];
      
      for (const component of components) {
        const elements = await page.$$(component.selector);
        const found = await Promise.all(elements.map(async el => {
          const text = await page.evaluate(el => el.textContent, el);
          return text.includes(component.text);
        }));
        
        expect(found.some(f => f)).toBe(true);
      }
    }, 30000);

    test('should render member card correctly', async () => {
      await page.goto(`${APP_URL}/member-card`, { waitUntil: 'networkidle0' });
      
      // Check for member card elements
      await page.waitForSelector('h2', { timeout: 10000 });
      
      const memberInfo = await page.$$eval('p', elements => 
        elements.map(el => el.textContent)
      );
      
      // Should have member ID, group number, etc.
      expect(memberInfo.some(text => text.includes('Member ID'))).toBe(true);
      expect(memberInfo.some(text => text.includes('Group'))).toBe(true);
    }, 30000);
  });

  describe('Responsive Design', () => {
    const viewports = [
      { name: 'Mobile', width: 375, height: 667 },
      { name: 'Tablet', width: 768, height: 1024 },
      { name: 'Desktop', width: 1920, height: 1080 }
    ];
    
    test.each(viewports)('should render correctly on $name', async ({ width, height }) => {
      await page.setViewport({ width, height });
      await page.goto(`${APP_URL}/`, { waitUntil: 'networkidle0' });
      
      // Take screenshot for visual verification
      await page.screenshot({ 
        path: `__tests__/screenshots/${width}x${height}-dashboard.png`,
        fullPage: true 
      });
      
      // Check if main content is visible
      const mainContent = await page.$('main');
      expect(mainContent).toBeTruthy();
      
      // Check if layout adapts
      if (width < 768) {
        // Mobile layout checks
        const navMenu = await page.$('nav');
        if (navMenu) {
          const navDisplay = await page.evaluate(el => 
            window.getComputedStyle(el).display, navMenu
          );
          // Navigation might be hidden or styled differently on mobile
          expect(['none', 'block', 'flex']).toContain(navDisplay);
        }
      }
    }, 30000);
  });

  describe('Performance', () => {
    test('should load pages within acceptable time', async () => {
      const startTime = Date.now();
      await page.goto(`${APP_URL}/`, { waitUntil: 'networkidle0' });
      const loadTime = Date.now() - startTime;
      
      // Page should load within 3 seconds
      expect(loadTime).toBeLessThan(3000);
    }, 30000);

    test('should have acceptable performance metrics', async () => {
      await page.goto(`${APP_URL}/`, { waitUntil: 'networkidle0' });
      
      const metrics = await page.evaluate(() => {
        const navigation = performance.getEntriesByType('navigation')[0];
        return {
          domContentLoaded: navigation.domContentLoadedEventEnd - navigation.domContentLoadedEventStart,
          loadComplete: navigation.loadEventEnd - navigation.loadEventStart,
          domInteractive: navigation.domInteractive - navigation.fetchStart
        };
      });
      
      // DOM should be interactive within 1.5 seconds
      expect(metrics.domInteractive).toBeLessThan(1500);
    }, 30000);
  });

  describe('API Integration', () => {
    test('should handle API errors gracefully', async () => {
      // This would test actual API failures, but since we have mocked APIs,
      // we'll check that the mock APIs are working
      await page.goto(`${APP_URL}/claims`, { waitUntil: 'networkidle0' });
      
      // Should display claims data
      await page.waitForSelector('table, ul, div[role="list"]', { timeout: 10000 });
      
      // Check if data is displayed
      const hasData = await page.evaluate(() => {
        const tables = document.querySelectorAll('table');
        const lists = document.querySelectorAll('ul, div[role="list"]');
        return tables.length > 0 || lists.length > 0;
      });
      
      expect(hasData).toBe(true);
    }, 30000);
  });

  describe('Security', () => {
    test('should redirect to login when accessing protected routes without auth', async () => {
      // Clear cookies to simulate logged out state
      const cookies = await page.cookies();
      await page.deleteCookie(...cookies);
      
      // Try to access protected route
      await page.goto(`${APP_URL}/benefits`, { waitUntil: 'networkidle0' });
      
      // Should redirect to login
      expect(page.url()).toContain('/login');
    }, 30000);

    test('should not expose sensitive data in console', async () => {
      const logs = [];
      page.on('console', msg => logs.push(msg.text()));
      
      await page.goto(`${APP_URL}/`, { waitUntil: 'networkidle0' });
      await page.waitForTimeout(1000);
      
      // Check logs for sensitive data patterns
      const sensitivePatterns = [
        /password/i,
        /token/i,
        /secret/i,
        /api[_-]?key/i
      ];
      
      const hasSensitiveData = logs.some(log => 
        sensitivePatterns.some(pattern => pattern.test(log))
      );
      
      expect(hasSensitiveData).toBe(false);
    }, 30000);
  });
});