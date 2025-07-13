import App from 'fusion-react';
import Router from 'fusion-plugin-react-router';
import Styletron from 'fusion-plugin-styletron-react';
import UniversalEvents from 'fusion-plugin-universal-events';
import PerformanceEmitter from 'fusion-plugin-browser-performance-emitter';
import RPC from 'fusion-plugin-rpc';
import I18n from 'fusion-plugin-i18n-react';
import ErrorHandling from 'fusion-plugin-error-handling';
import CsrfProtection from 'fusion-plugin-csrf-protection';
import JWTPlugin from 'fusion-plugin-jwt';

import root from './root.js';
import api from './services/api.js';

export default () => {
  const app = new App(root);

  // Register plugins
  app.register(Styletron);
  app.register(Router);
  
  // Security plugins
  app.register(CsrfProtection);
  app.register(JWTPlugin);
  
  // RPC for API communication
  app.register(RPC, {
    handlers: api.handlers,
  });
  
  // Performance monitoring
  if (__BROWSER__) {
    app.register(UniversalEvents);
    app.register(PerformanceEmitter);
  }
  
  // Error handling
  app.register(ErrorHandling);
  
  // Internationalization
  app.register(I18n, {
    translations: {
      en_US: require('./translations/en_US.json'),
    },
  });

  return app;
};