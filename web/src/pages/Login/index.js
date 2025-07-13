import React, {useState} from 'react';
import {styled} from 'fusion-plugin-styletron-react';
import {useHistory} from 'fusion-plugin-react-router';

const LoginContainer = styled('div', {
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  minHeight: '100vh',
  backgroundColor: '#f5f5f5',
});

const LoginCard = styled('div', {
  backgroundColor: '#fff',
  borderRadius: '8px',
  boxShadow: '0 4px 12px rgba(0,0,0,0.1)',
  padding: '40px',
  width: '100%',
  maxWidth: '400px',
});

const Logo = styled('h1', {
  fontSize: '32px',
  fontWeight: 'bold',
  textAlign: 'center',
  marginBottom: '8px',
  color: '#000',
});

const Subtitle = styled('p', {
  fontSize: '16px',
  color: '#666',
  textAlign: 'center',
  marginBottom: '32px',
});

const Form = styled('form', {
  display: 'flex',
  flexDirection: 'column',
  gap: '20px',
});

const FormGroup = styled('div', {
  display: 'flex',
  flexDirection: 'column',
  gap: '8px',
});

const Label = styled('label', {
  fontSize: '14px',
  fontWeight: '500',
  color: '#333',
});

const Input = styled('input', {
  padding: '12px 16px',
  border: '1px solid #ddd',
  borderRadius: '6px',
  fontSize: '16px',
  outline: 'none',
  transition: 'border-color 0.2s',
  ':focus': {
    borderColor: '#276ef1',
  },
});

const Button = styled('button', ({$variant}) => ({
  padding: '12px 24px',
  borderRadius: '6px',
  fontSize: '16px',
  fontWeight: '500',
  cursor: 'pointer',
  transition: 'all 0.2s',
  border: 'none',
  backgroundColor: $variant === 'primary' ? '#276ef1' : '#f0f0f0',
  color: $variant === 'primary' ? '#fff' : '#333',
  ':hover': {
    backgroundColor: $variant === 'primary' ? '#1a5ed8' : '#e0e0e0',
  },
  ':disabled': {
    opacity: 0.6,
    cursor: 'not-allowed',
  },
}));

const ErrorMessage = styled('div', {
  backgroundColor: '#fce8e6',
  color: '#d32f2f',
  padding: '12px',
  borderRadius: '6px',
  fontSize: '14px',
});

const DemoInfo = styled('div', {
  marginTop: '24px',
  padding: '16px',
  backgroundColor: '#e8f4fd',
  borderRadius: '6px',
  fontSize: '14px',
  color: '#1976d2',
});

const Login = () => {
  const history = useHistory();
  const [formData, setFormData] = useState({
    memberId: '',
    password: '',
  });
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
    setError('');
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');

    // Mock authentication
    setTimeout(() => {
      if (formData.memberId === 'M123456' && formData.password === 'demo') {
        // Store mock auth token
        localStorage.setItem('authToken', 'mock-jwt-token');
        history.push('/');
      } else {
        setError('Invalid member ID or password');
      }
      setLoading(false);
    }, 1000);
  };

  return (
    <LoginContainer>
      <LoginCard>
        <Logo>Sydney Health</Logo>
        <Subtitle>Sign in to manage your health benefits</Subtitle>
        
        <Form onSubmit={handleSubmit}>
          {error && <ErrorMessage>{error}</ErrorMessage>}
          
          <FormGroup>
            <Label htmlFor="memberId">Member ID</Label>
            <Input
              id="memberId"
              name="memberId"
              type="text"
              placeholder="Enter your member ID"
              value={formData.memberId}
              onChange={handleChange}
              required
            />
          </FormGroup>
          
          <FormGroup>
            <Label htmlFor="password">Password</Label>
            <Input
              id="password"
              name="password"
              type="password"
              placeholder="Enter your password"
              value={formData.password}
              onChange={handleChange}
              required
            />
          </FormGroup>
          
          <Button
            type="submit"
            $variant="primary"
            disabled={loading}
          >
            {loading ? 'Signing in...' : 'Sign In'}
          </Button>
        </Form>
        
        <DemoInfo>
          <strong>Demo Credentials:</strong><br />
          Member ID: M123456<br />
          Password: demo
        </DemoInfo>
      </LoginCard>
    </LoginContainer>
  );
};

export default Login;