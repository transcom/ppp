import { combineReducers } from 'redux';
import { reducer as formReducer } from 'redux-form';
import { connectRouter } from 'connected-react-router';
import { adminReducer } from 'react-admin';
import defaultMessages from 'ra-language-english';
import storage from 'redux-persist/lib/storage';
import { persistReducer } from 'redux-persist';

import authReducer from 'store/auth/reducer';
import onboardingReducer from 'store/onboarding/reducer';
import flashReducer from 'store/flash/reducer';
import userReducer from 'shared/Data/users';
import { swaggerReducerPublic, swaggerReducerInternal } from 'shared/Swagger/ducks';
import { requestsReducer } from 'shared/Swagger/requestsReducer';
import { entitiesReducer } from 'shared/Entities/reducer';
import uiReducer from 'shared/UI/ducks';
import { moveReducer } from 'scenes/Moves/ducks';
import { ppmReducer } from 'scenes/Moves/Ppm/ducks';
import { serviceMemberReducer } from 'scenes/ServiceMembers/ducks';
import { ordersReducer } from 'scenes/Orders/ducks';
import { signedCertificationReducer } from 'scenes/Legalese/ducks';
import { reviewReducer } from 'scenes/Review/ducks';
import { officeFlashMessagesReducer } from 'scenes/Office/ducks';
import officePpmReducer from 'scenes/Office/Ppm/ducks';

const locale = 'en';
const i18nProvider = () => defaultMessages;

const authPersistConfig = {
  key: 'auth',
  storage,
  whitelist: ['activeRole'],
};

const defaultReducers = {
  auth: persistReducer(authPersistConfig, authReducer),
  flash: flashReducer,
  form: formReducer,
  swaggerPublic: swaggerReducerPublic,
  requests: requestsReducer,
  ui: uiReducer,
  user: userReducer,
  entities: entitiesReducer,
};

export const appReducer = (history) =>
  combineReducers({
    ...defaultReducers,
    onboarding: onboardingReducer,
    router: connectRouter(history),
    swaggerInternal: swaggerReducerInternal,
    moves: moveReducer,
    ppm: ppmReducer,
    serviceMember: serviceMemberReducer,
    orders: ordersReducer,
    signedCertification: signedCertificationReducer,
    review: reviewReducer,
    flashMessages: officeFlashMessagesReducer,
    ppmIncentive: officePpmReducer,
  });

export const adminAppReducer = (history) =>
  combineReducers({
    ...defaultReducers,
    admin: adminReducer,
    i18n: i18nProvider(locale),
    form: formReducer,
    router: connectRouter(history),
  });

export default appReducer;
