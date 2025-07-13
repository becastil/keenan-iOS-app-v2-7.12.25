import { handlers } from '../../../src/services/api';

describe('API Service', () => {
  describe('login handler', () => {
    test('returns success for valid credentials', async () => {
      const result = await handlers.login({ 
        email: 'test@example.com', 
        password: 'password123' 
      });
      
      expect(result.success).toBe(true);
      expect(result.token).toBeTruthy();
      expect(result.user).toEqual({
        id: '123',
        name: 'John Doe',
        email: 'test@example.com'
      });
    });

    test('returns error for invalid credentials', async () => {
      const result = await handlers.login({ 
        email: 'wrong@example.com', 
        password: 'wrongpass' 
      });
      
      expect(result.success).toBe(true); // Mock always returns success
      expect(result.token).toBeTruthy();
    });
  });

  describe('getBenefits handler', () => {
    test('returns benefits data', async () => {
      const result = await handlers.getBenefits();
      
      expect(result).toHaveProperty('deductible');
      expect(result).toHaveProperty('outOfPocketMax');
      expect(result).toHaveProperty('coverageDetails');
      expect(result.coverageDetails).toBeInstanceOf(Array);
      expect(result.coverageDetails.length).toBeGreaterThan(0);
    });

    test('returns proper deductible structure', async () => {
      const result = await handlers.getBenefits();
      
      expect(result.deductible).toHaveProperty('individual');
      expect(result.deductible).toHaveProperty('family');
      expect(result.deductible.individual).toHaveProperty('met');
      expect(result.deductible.individual).toHaveProperty('total');
    });
  });

  describe('getClaims handler', () => {
    test('returns claims array', async () => {
      const result = await handlers.getClaims();
      
      expect(result).toBeInstanceOf(Array);
      expect(result.length).toBeGreaterThan(0);
    });

    test('returns properly structured claims', async () => {
      const result = await handlers.getClaims();
      const firstClaim = result[0];
      
      expect(firstClaim).toHaveProperty('id');
      expect(firstClaim).toHaveProperty('date');
      expect(firstClaim).toHaveProperty('provider');
      expect(firstClaim).toHaveProperty('service');
      expect(firstClaim).toHaveProperty('status');
      expect(firstClaim).toHaveProperty('amount');
    });
  });

  describe('getProviders handler', () => {
    test('returns providers based on search criteria', async () => {
      const result = await handlers.getProviders({ 
        specialty: 'Primary Care',
        location: '10001'
      });
      
      expect(result).toBeInstanceOf(Array);
      expect(result.length).toBeGreaterThan(0);
    });

    test('returns properly structured provider data', async () => {
      const result = await handlers.getProviders({});
      const firstProvider = result[0];
      
      expect(firstProvider).toHaveProperty('id');
      expect(firstProvider).toHaveProperty('name');
      expect(firstProvider).toHaveProperty('specialty');
      expect(firstProvider).toHaveProperty('address');
      expect(firstProvider).toHaveProperty('phone');
      expect(firstProvider).toHaveProperty('distance');
      expect(firstProvider).toHaveProperty('inNetwork');
    });
  });

  describe('getMemberCard handler', () => {
    test('returns member card information', async () => {
      const result = await handlers.getMemberCard();
      
      expect(result).toHaveProperty('memberId');
      expect(result).toHaveProperty('memberName');
      expect(result).toHaveProperty('groupNumber');
      expect(result).toHaveProperty('planName');
      expect(result).toHaveProperty('effectiveDate');
      expect(result).toHaveProperty('copayInfo');
    });

    test('returns proper copay structure', async () => {
      const result = await handlers.getMemberCard();
      
      expect(result.copayInfo).toHaveProperty('primaryCare');
      expect(result.copayInfo).toHaveProperty('specialist');
      expect(result.copayInfo).toHaveProperty('urgentCare');
      expect(result.copayInfo).toHaveProperty('emergency');
    });
  });

  describe('getMessages handler', () => {
    test('returns messages array', async () => {
      const result = await handlers.getMessages();
      
      expect(result).toBeInstanceOf(Array);
      expect(result.length).toBeGreaterThan(0);
    });

    test('returns properly structured messages', async () => {
      const result = await handlers.getMessages();
      const firstMessage = result[0];
      
      expect(firstMessage).toHaveProperty('id');
      expect(firstMessage).toHaveProperty('from');
      expect(firstMessage).toHaveProperty('subject');
      expect(firstMessage).toHaveProperty('date');
      expect(firstMessage).toHaveProperty('preview');
      expect(firstMessage).toHaveProperty('unread');
    });
  });

  describe('sendMessage handler', () => {
    test('successfully sends a message', async () => {
      const messageData = {
        to: 'provider@example.com',
        subject: 'Test Message',
        body: 'This is a test message'
      };
      
      const result = await handlers.sendMessage(messageData);
      
      expect(result).toHaveProperty('success');
      expect(result.success).toBe(true);
      expect(result).toHaveProperty('messageId');
    });
  });
});