# Sydney Health Web App - Manual Testing Guide

## Quick Start Testing Instructions

Since automated E2E testing is encountering environment issues (Puppeteer dependencies in WSL), please follow this manual testing guide:

### 1. Start the Application

```bash
cd web/
npm run dev
```

The application should start on `http://localhost:3000`

### 2. Login Testing

- Open browser: `http://localhost:3000`
- You should be redirected to `/login`
- Login credentials:
  - Email: `test@example.com`
  - Password: `password123`
- Alternative (if member ID/PIN login exists):
  - Member ID: `M123456`
  - PIN: `demo`

### 3. Page Navigation Checklist

After login, verify each page loads correctly:

- [ ] **Dashboard** (`/`)
  - [ ] Deductible Tracker widget displays
  - [ ] Recent Claims widget shows sample claims
  - [ ] Quick Actions buttons are clickable
  
- [ ] **Benefits** (`/benefits`)
  - [ ] Coverage details table displays
  - [ ] Deductible information shows
  - [ ] Out-of-pocket maximum displays

- [ ] **Claims** (`/claims`)
  - [ ] Claims list/table renders
  - [ ] Claims show date, provider, amount, status
  - [ ] Filter or search functionality (if implemented)

- [ ] **Providers** (`/providers`)
  - [ ] Search form displays
  - [ ] Provider list shows after search
  - [ ] In-network indicator visible

- [ ] **Member Card** (`/member-card`)
  - [ ] Digital card displays
  - [ ] Member ID visible
  - [ ] Group number shows
  - [ ] Copay information listed

- [ ] **Messages** (`/messages`)
  - [ ] Message list displays
  - [ ] Unread indicator works
  - [ ] Message compose button (if exists)

### 4. Responsive Design Testing

Test on different screen sizes:

1. **Desktop** (1920x1080)
   - F12 → Responsive mode → Select Desktop
   - Navigation should be visible on side
   - Content should use full width

2. **Tablet** (768x1024)
   - F12 → iPad or similar device
   - Layout should adapt appropriately
   - Navigation may collapse

3. **Mobile** (375x667)
   - F12 → iPhone or similar device
   - Navigation should be mobile-friendly
   - Content should stack vertically

### 5. Console Error Check

- Open Developer Tools (F12)
- Go to Console tab
- Navigate through all pages
- **No red errors should appear**
- Warnings (yellow) are acceptable

### 6. Performance Quick Check

- Open Network tab in Developer Tools
- Refresh page
- Check:
  - Page load time < 3 seconds
  - Bundle size reasonable (< 1-2MB)
  - No failed requests (red items)

## Expected Results

✅ **PASS Criteria:**
- All pages load without errors
- Navigation works correctly
- Login/logout flow functions
- No console errors
- Responsive design adapts properly
- Performance is acceptable

❌ **FAIL Criteria:**
- Pages show blank screens
- Navigation broken
- Console shows red errors
- Login doesn't work
- Mobile view unusable

## Common Issues & Solutions

### Fusion CLI Not Found
```bash
# Install globally
npm install -g fusion-cli

# Or use npx (already configured)
npm run dev
```

### Port 3000 Already in Use
```bash
# Find process using port
lsof -i :3000
# Kill process
kill -9 [PID]
```

### Login Not Working
- Check browser console for errors
- Verify mock API is returning data
- Clear browser cache/cookies

## Report Your Findings

After testing, note:
1. Which tests passed/failed
2. Any console errors encountered
3. Performance issues noticed
4. UI/UX problems found
5. Browser/device tested on

This manual testing covers all the critical paths that the automated E2E tests would have validated.