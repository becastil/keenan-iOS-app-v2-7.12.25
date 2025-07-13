# Sydney Health iOS App

iOS application built with Swift and RIBs architecture, replicating the functionality of Anthem Sydney Health app.

## Architecture

This app uses Uber's RIBs (Router, Interactor, Builder) architecture pattern for:
- Clear separation of concerns
- Testability
- Deep linking support
- Business logic isolation

## Tech Stack

- **Language**: Swift 5.9+
- **Architecture**: RIBs
- **UI**: UIKit + SnapKit
- **Reactive**: RxSwift + RxCocoa
- **Networking**: Alamofire + SwiftGRPC
- **Image Loading**: Kingfisher
- **Dependencies**: CocoaPods

## Requirements

- iOS 13.0+
- Xcode 15.0+
- CocoaPods 1.12+

## Setup

1. Install dependencies:
```bash
cd ios
pod install
```

2. Open the workspace:
```bash
open SydneyHealth.xcworkspace
```

3. Build and run the project in Xcode

## Project Structure

```
SydneyHealth/
├── AppStart/           # App launch and configuration
├── RIBs/              # RIBs components
│   ├── Root/          # Root RIB
│   ├── LoggedOut/     # Login flow
│   ├── LoggedIn/      # Main app flow
│   ├── Dashboard/     # Dashboard RIB
│   ├── Benefits/      # Benefits RIB
│   ├── Claims/        # Claims RIB
│   ├── Providers/     # Provider search RIB
│   ├── MemberCard/    # Member ID RIB
│   └── Messages/      # Messaging RIB
├── Services/          # Business logic services
├── Models/            # Data models
├── Utils/             # Utilities and extensions
└── Resources/         # Assets and resources
```

## RIBs Architecture

Each RIB consists of:
- **Router**: Manages child RIBs and navigation
- **Interactor**: Contains business logic
- **Builder**: Creates and wires up the RIB
- **View**: UI components (optional)
- **Presenter**: Formats data for display (optional)

## Key Features

- Biometric authentication (Face ID/Touch ID)
- Offline support with local caching
- Push notifications for claims updates
- Deep linking support
- Real-time messaging with healthcare providers

## Testing

Run unit tests:
```bash
xcodebuild test -workspace SydneyHealth.xcworkspace -scheme SydneyHealth -destination 'platform=iOS Simulator,name=iPhone 15'
```

## Code Style

We use SwiftLint for code style enforcement. Rules are defined in `.swiftlint.yml`.

## Authentication

The app uses JWT tokens for authentication, stored securely in the iOS Keychain.

Demo credentials:
- Member ID: M123456
- Password: demo