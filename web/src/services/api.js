import {createRPCHandler} from 'fusion-plugin-rpc';

// Mock API handlers
const handlers = createRPCHandler({
  // Member endpoints
  getMember: async (args, ctx) => {
    // Mock member data
    return {
      memberId: 'M123456',
      firstName: 'John',
      lastName: 'Doe',
      email: 'john.doe@email.com',
      phone: '+1-555-123-4567',
      dateOfBirth: '1985-06-15',
      address: {
        street1: '123 Main Street',
        street2: 'Apt 4B',
        city: 'San Francisco',
        state: 'CA',
        zipCode: '94105',
      },
    };
  },

  // Benefits endpoints
  getBenefits: async (args, ctx) => {
    return [
      {
        id: 'BEN001',
        name: 'Primary Care Visit',
        coverageType: 'medical',
        copay: '$20',
        coinsurance: '0%',
      },
      {
        id: 'BEN002',
        name: 'Specialist Visit',
        coverageType: 'medical',
        copay: '$40',
        coinsurance: '0%',
      },
    ];
  },

  // Claims endpoints
  getClaims: async (args, ctx) => {
    return {
      claims: [
        {
          id: 'CLM001234',
          date: '2024-01-15',
          provider: 'Bay Area Medical Center',
          service: 'Office Visit',
          amount: 250.00,
          yourCost: 25.00,
          status: 'Paid',
        },
      ],
      totalCount: 1,
    };
  },

  // Provider endpoints
  searchProviders: async (args, ctx) => {
    const {specialty, location, radius} = args;
    
    return {
      providers: [
        {
          id: 'PRV001',
          name: 'Dr. Sarah Johnson',
          specialty: specialty || 'Primary Care',
          address: '456 Market St, San Francisco, CA',
          distance: 2.5,
          rating: 4.8,
          acceptingNewPatients: true,
        },
      ],
      totalCount: 1,
    };
  },

  // Messaging endpoints
  getConversations: async (args, ctx) => {
    return {
      conversations: [
        {
          id: 'CONV001',
          subject: 'Question about benefits',
          lastMessage: 'Thank you for your response',
          unreadCount: 0,
          updatedAt: '2024-01-15T10:30:00Z',
        },
      ],
      totalCount: 1,
    };
  },

  sendMessage: async (args, ctx) => {
    const {conversationId, content} = args;
    
    return {
      id: 'MSG001',
      conversationId,
      content,
      sentAt: new Date().toISOString(),
    };
  },
});

export default {handlers};