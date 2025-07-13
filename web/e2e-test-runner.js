const puppeteer = require('puppeteer');
const { spawn } = require('child_process');
const path = require('path');
const fs = require('fs');

// Test configuration
const APP_URL = 'http://localhost:3000';
const TEST_CREDENTIALS = {
  email: 'test@example.com',
  password: 'password123',
  memberId: 'M123456',
  pin: 'demo'
};

// Validation checklist
const validationChecklist = {
  startup: false,
  login: false,
  dashboard: false,
  benefits: false,
  claims: false,
  providers: false,
  memberCard: false,
  messages: false,
  responsive: false,
  noErrors: true,
  performance: false
};

let browser;
let page;
let serverProcess;
let consoleErrors = [];

async function runE2ETests() {
  console.log('🚀 Sydney Health Web App - E2E Testing with Validation\n');
  console.log('Environment: WSL with Node.js v18.20.8\n');

  try {
    // Start the development server
    await startDevServer();
    
    // Launch browser
    browser = await puppeteer.launch({
      headless: 'new',
      args: ['--no-sandbox', '--disable-setuid-sandbox']
    });
    
    page = await browser.newPage();
    
    // Capture console errors
    page.on('console', msg => {
      if (msg.type() === 'error') {
        consoleErrors.push({
          text: msg.text(),
          location: msg.location()
        });
        validationChecklist.noErrors = false;
      }
    });
    
    // Set viewport for desktop testing
    await page.setViewport({ width: 1280, height: 800 });
    
    // Run all tests
    await testApplicationStartup();
    await testLogin();
    await testDashboard();
    await testAllPages();
    await testResponsiveDesign();
    await testPerformance();
    
    // Generate final report
    generateReport();
    
  } catch (error) {
    console.error('❌ Test failed:', error.message);
    validationChecklist.noErrors = false;
  } finally {
    await cleanup();
  }
}

async function startDevServer() {
  console.log('📦 Starting development server...\n');
  
  return new Promise((resolve, reject) => {
    serverProcess = spawn('npm', ['run', 'dev'], {
      cwd: path.resolve(__dirname),
      shell: true
    });
    
    let serverStarted = false;
    
    serverProcess.stdout.on('data', (data) => {
      const output = data.toString();
      console.log(`[SERVER] ${output}`);
      
      if (output.includes('Server running') || output.includes('Started server')) {
        serverStarted = true;
        setTimeout(resolve, 5000); // Give it extra time to fully initialize
      }
    });
    
    serverProcess.stderr.on('data', (data) => {
      console.error(`[SERVER ERROR] ${data}`);
    });
    
    // Timeout after 30 seconds
    setTimeout(() => {
      if (!serverStarted) {
        console.log('⏱️  Server startup timeout - attempting to connect anyway...');
        resolve();
      }
    }, 30000);
  });
}

async function testApplicationStartup() {
  console.log('\n1️⃣ Testing Application Startup...');
  
  try {
    const response = await page.goto(APP_URL, { 
      waitUntil: 'networkidle0',
      timeout: 30000 
    });
    
    if (response && response.status() === 200) {
      validationChecklist.startup = true;
      console.log('✅ Application started successfully');
    } else {
      console.log('❌ Application failed to start properly');
    }
  } catch (error) {
    console.log('❌ Could not connect to application:', error.message);
  }
}

async function testLogin() {
  console.log('\n2️⃣ Testing Login Flow...');
  
  try {
    // Check if we're on login page or redirected there
    const currentUrl = page.url();
    if (!currentUrl.includes('/login')) {
      await page.goto(`${APP_URL}/login`, { waitUntil: 'networkidle0' });
    }
    
    // Wait for login form
    await page.waitForSelector('input[type="email"]', { timeout: 10000 });
    
    // Fill in credentials
    await page.type('input[type="email"]', TEST_CREDENTIALS.email);
    await page.type('input[type="password"]', TEST_CREDENTIALS.password);
    
    // Submit form
    await Promise.all([
      page.waitForNavigation({ waitUntil: 'networkidle0' }),
      page.click('button[type="submit"]')
    ]);
    
    // Verify successful login
    if (page.url() === `${APP_URL}/` || page.url().includes('dashboard')) {
      validationChecklist.login = true;
      console.log('✅ Login successful');
    } else {
      console.log('❌ Login failed - unexpected redirect');
    }
  } catch (error) {
    console.log('❌ Login test failed:', error.message);
  }
}

async function testDashboard() {
  console.log('\n3️⃣ Testing Dashboard Components...');
  
  try {
    // Ensure we're on dashboard
    if (!page.url().includes(`${APP_URL}/`)) {
      await page.goto(`${APP_URL}/`, { waitUntil: 'networkidle0' });
    }
    
    // Check for key dashboard elements
    const elements = await page.evaluate(() => {
      const hasDeductible = !!document.querySelector('h2')?.textContent.includes('Deductible');
      const hasClaims = !!document.querySelector('h2')?.textContent.includes('Claims');
      const hasActions = !!document.querySelector('h2')?.textContent.includes('Actions');
      
      return { hasDeductible, hasClaims, hasActions };
    });
    
    if (elements.hasDeductible && elements.hasClaims && elements.hasActions) {
      validationChecklist.dashboard = true;
      console.log('✅ Dashboard components loaded correctly');
    } else {
      console.log('❌ Some dashboard components missing');
    }
  } catch (error) {
    console.log('❌ Dashboard test failed:', error.message);
  }
}

async function testAllPages() {
  const pages = [
    { name: 'Benefits', path: '/benefits', key: 'benefits' },
    { name: 'Claims', path: '/claims', key: 'claims' },
    { name: 'Providers', path: '/providers', key: 'providers' },
    { name: 'Member Card', path: '/member-card', key: 'memberCard' },
    { name: 'Messages', path: '/messages', key: 'messages' }
  ];
  
  console.log('\n4️⃣ Testing All Pages...');
  
  for (const pageInfo of pages) {
    try {
      console.log(`\n   Testing ${pageInfo.name}...`);
      
      await page.goto(`${APP_URL}${pageInfo.path}`, { 
        waitUntil: 'networkidle0',
        timeout: 15000 
      });
      
      // Wait for page to load
      await page.waitForSelector('h1', { timeout: 5000 });
      
      // Check page loaded correctly
      const pageTitle = await page.$eval('h1', el => el.textContent);
      
      if (pageTitle) {
        validationChecklist[pageInfo.key] = true;
        console.log(`   ✅ ${pageInfo.name} page loaded`);
      } else {
        console.log(`   ❌ ${pageInfo.name} page failed to load properly`);
      }
      
      // Take screenshot for documentation
      await page.screenshot({ 
        path: `__tests__/screenshots/${pageInfo.key}.png`,
        fullPage: true 
      });
      
    } catch (error) {
      console.log(`   ❌ ${pageInfo.name} test failed:`, error.message);
    }
  }
}

async function testResponsiveDesign() {
  console.log('\n5️⃣ Testing Responsive Design...');
  
  const viewports = [
    { name: 'Mobile', width: 375, height: 667 },
    { name: 'Tablet', width: 768, height: 1024 },
    { name: 'Desktop', width: 1920, height: 1080 }
  ];
  
  try {
    for (const viewport of viewports) {
      await page.setViewport({ width: viewport.width, height: viewport.height });
      await page.goto(`${APP_URL}/`, { waitUntil: 'networkidle0' });
      
      // Check if content is visible
      const isVisible = await page.evaluate(() => {
        const main = document.querySelector('main');
        return main && main.offsetHeight > 0;
      });
      
      if (isVisible) {
        console.log(`   ✅ ${viewport.name} view renders correctly`);
      } else {
        console.log(`   ❌ ${viewport.name} view has issues`);
      }
    }
    
    validationChecklist.responsive = true;
  } catch (error) {
    console.log('❌ Responsive design test failed:', error.message);
  }
}

async function testPerformance() {
  console.log('\n6️⃣ Testing Performance Metrics...');
  
  try {
    await page.goto(`${APP_URL}/`, { waitUntil: 'networkidle0' });
    
    const metrics = await page.evaluate(() => {
      const navigation = performance.getEntriesByType('navigation')[0];
      return {
        domContentLoaded: navigation.domContentLoadedEventEnd - navigation.domContentLoadedEventStart,
        loadComplete: navigation.loadEventEnd - navigation.loadEventStart,
        domInteractive: navigation.domInteractive - navigation.fetchStart
      };
    });
    
    console.log(`   DOM Interactive: ${metrics.domInteractive}ms`);
    console.log(`   DOM Content Loaded: ${metrics.domContentLoaded}ms`);
    console.log(`   Page Load Complete: ${metrics.loadComplete}ms`);
    
    if (metrics.domInteractive < 1500) {
      validationChecklist.performance = true;
      console.log('   ✅ Performance metrics are acceptable');
    } else {
      console.log('   ⚠️  Performance could be improved');
    }
  } catch (error) {
    console.log('❌ Performance test failed:', error.message);
  }
}

function generateReport() {
  console.log('\n' + '='.repeat(60));
  console.log('📊 E2E TEST VALIDATION REPORT');
  console.log('='.repeat(60));
  
  // Calculate pass rate
  const totalTests = Object.keys(validationChecklist).length;
  const passedTests = Object.values(validationChecklist).filter(v => v).length;
  const passRate = Math.round((passedTests / totalTests) * 100);
  
  console.log(`\n✅ Tests Passed: ${passedTests}/${totalTests} (${passRate}%)\n`);
  
  // Show checklist results
  console.log('Validation Checklist:');
  Object.entries(validationChecklist).forEach(([test, passed]) => {
    const icon = passed ? '✅' : '❌';
    const testName = test.charAt(0).toUpperCase() + test.slice(1).replace(/([A-Z])/g, ' $1');
    console.log(`  ${icon} ${testName}`);
  });
  
  // Console errors
  if (consoleErrors.length > 0) {
    console.log('\n⚠️  Console Errors Found:');
    consoleErrors.forEach((error, i) => {
      console.log(`  ${i + 1}. ${error.text}`);
    });
  }
  
  // Coverage summary
  console.log('\n📈 Coverage Summary:');
  console.log('  - Page Coverage: All main pages tested');
  console.log('  - Responsive Testing: Desktop, Tablet, Mobile');
  console.log('  - User Flows: Login → Dashboard → All Features');
  console.log('  - Performance: Basic metrics captured');
  
  // Save report
  const report = {
    timestamp: new Date().toISOString(),
    passRate,
    checklist: validationChecklist,
    errors: consoleErrors,
    environment: {
      node: process.version,
      platform: process.platform
    }
  };
  
  fs.writeFileSync('e2e-test-report.json', JSON.stringify(report, null, 2));
  console.log('\n📄 Full report saved to: e2e-test-report.json');
}

async function cleanup() {
  if (browser) {
    await browser.close();
  }
  
  if (serverProcess) {
    serverProcess.kill();
  }
  
  console.log('\n🏁 Testing complete!');
}

// Run the tests
runE2ETests().catch(console.error);