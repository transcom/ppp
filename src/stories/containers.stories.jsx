import React from 'react';
import { storiesOf } from '@storybook/react';

// Containers

storiesOf('Components|Containers', module).add('all', () => (
  <div id="containers" style={{ padding: '20px' }}>
    <div className="container">
      <code>
        <b>Container Default</b>
        <br />
        .container
      </code>
    </div>
    <div className="container container--gray">
      <code>
        <b>Container Gray</b>
        <br />
        .container
        <br />
        .container--gray
      </code>
    </div>
    <div className="container container--popout">
      <code>
        <b>Container Popout</b>
        <br />
        .container
        <br />
        .container--popout
      </code>
    </div>
    <div className="container container--accent--blue">
      <code>
        <b>Container Accent Blue</b>
        <br />
        .container
        <br />
        .container--accent--blue
      </code>
    </div>
    <div className="container container--accent--yellow">
      <code>
        <b>Container Accent Yellow</b>
        <br />
        .container
        <br />
        .container--accent--yellow
      </code>
    </div>
    <div className="container container--accent--hhg">
      <code>
        <b>Container Accent HHG</b>
        <br />
        .container
        <br />
        .container--accent--hhg
      </code>
    </div>
    <div className="container container--accent--ppm">
      <code>
        <b>Container Accent PPM</b>
        <br />
        .container
        <br />
        .container--accent--ppm
      </code>
    </div>
    <div className="container container--accent--ub">
      <code>
        <b>Container Accent UB</b>
        <br />
        .container
        <br />
        .container--accent--ub
      </code>
    </div>
    <div className="container container--accent--nts">
      <code>
        <b>Container Accent NTS</b>
        <br />
        .container
        <br />
        .container--accent--nts
      </code>
    </div>
    <div className="container container--accent--ntsr">
      <code>
        <b>Container Accent NTSR</b>
        <br />
        .container
        <br />
        .container--accent--ntsr
      </code>
    </div>
  </div>
));
