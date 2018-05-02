import React, { Component } from 'react';
import { Redirect, Route, Switch } from 'react-router-dom';
import { ConnectedRouter } from 'react-router-redux';
import { history } from 'shared/store';

import QueueHeader from 'shared/Header/Office';
import QueueList from './QueueList';
import QueueTable from './QueueTable';

class Queues extends Component {
  render() {
    return (
      <div className="usa-grid">
        <div className="usa-width-one-fourth">
          <QueueList />
        </div>
        <div className="usa-width-three-fourths">
          <QueueTable queueType={this.props.match.params.queueType} />
        </div>
      </div>
    );
  }
}

class OfficeWrapper extends Component {
  componentDidMount() {
    document.title = 'Transcom PPP: Office';
  }

  render() {
    return (
      <ConnectedRouter history={history}>
        <div className="Office site">
          <QueueHeader />
          <main className="site__content">
            <div>
              <div className="usa-grid" />
              <Switch>
                <Redirect from="/" to="/queues/new" exact />
                <Route path="/queues/:queueType" component={Queues} />
              </Switch>
            </div>
          </main>
        </div>
      </ConnectedRouter>
    );
  }
}

export default OfficeWrapper;
