# Sydney Health Web Application

A modern healthcare application built with Fusion.js, replicating the functionality of Anthem Sydney Health app.

## Features

- View medical, dental, and vision benefits
- Search for in-network providers
- Display digital member ID cards
- Review claims (pending and approved)
- Access cost estimates and care guidance
- Send and receive secure member messages

## Tech Stack

- **Framework**: Fusion.js (Uber's React framework)
- **UI Styling**: Styletron (CSS-in-JS)
- **Routing**: fusion-plugin-react-router
- **API Communication**: fusion-plugin-rpc
- **Authentication**: fusion-plugin-jwt
- **State Management**: React hooks and context

## Getting Started

### Prerequisites

- Node.js >= 16.0.0
- npm or yarn

### Installation

```bash
# From the web directory
npm install
```

### Development

```bash
# Start the development server
npm run dev
```

The app will be available at http://localhost:3000

### Build

```bash
# Build for production
npm run build

# Start production server
npm start
```

## Project Structure

```
web/
├── src/
│   ├── components/       # Reusable UI components
│   ├── pages/           # Page components
│   ├── services/        # API services
│   ├── styles/          # Global styles
│   ├── utils/           # Utility functions
│   ├── translations/    # i18n translations
│   ├── main.js         # Fusion.js app entry
│   └── root.js         # Root React component
├── package.json
└── README.md
```

## Key Components

- **Layout**: Main app layout with sidebar navigation
- **Dashboard**: Overview of member benefits and recent activity
- **Benefits**: Detailed benefit information by category
- **Claims**: View and manage healthcare claims
- **Providers**: Search for in-network healthcare providers
- **MemberCard**: Digital member ID cards
- **Messages**: Secure messaging with support

## Authentication

The app uses JWT tokens for authentication. Login with:
- Member ID: M123456
- Password: demo

## API Integration

The app communicates with backend microservices through the API Gateway using gRPC/HTTP bridge.

## Styling

We use Styletron for CSS-in-JS styling, following Uber's Base Design System patterns.

## Testing

```bash
npm test
```

## Deployment

The app is configured for deployment with Uber's infrastructure but can be adapted for other platforms.