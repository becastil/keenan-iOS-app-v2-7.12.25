import React from 'react';
import { render, screen } from '@testing-library/react';
import Card from '../../../src/components/Card';

// Mock styled from fusion-plugin-styletron-react
jest.mock('fusion-plugin-styletron-react', () => ({
  styled: (component) => (styles) => {
    const StyledComponent = ({ children, ...props }) => {
      return React.createElement(component, props, children);
    };
    return StyledComponent;
  }
}));

describe('Card Component', () => {
  test('renders with title', () => {
    render(<Card title="Test Card" />);
    expect(screen.getByText('Test Card')).toBeInTheDocument();
  });

  test('renders children content', () => {
    render(
      <Card title="Test Card">
        <p>Card content</p>
      </Card>
    );
    expect(screen.getByText('Card content')).toBeInTheDocument();
  });

  test('renders without title', () => {
    render(
      <Card>
        <p>Just content</p>
      </Card>
    );
    expect(screen.getByText('Just content')).toBeInTheDocument();
  });

  test('applies proper structure', () => {
    const { container } = render(
      <Card title="Structured Card">
        <div>Content</div>
      </Card>
    );
    
    // Check that card has proper nesting
    const cardElement = container.firstChild;
    expect(cardElement).toBeTruthy();
    expect(cardElement.tagName).toBe('DIV');
  });
});