import React from 'react';
import { string, node, bool } from 'prop-types';
import classnames from 'classnames';

import Hint from 'shared/Hint';

const Fieldset = ({ children, legend, className, legendSrOnly, legendClassName, hintText }) => {
  const classes = classnames('usa-fieldset', className);

  const legendClasses = classnames(`usa-legend legend-container ${legendClassName}`, {
    'usa-sr-only': legendSrOnly,
  });

  return (
    <fieldset data-testid="fieldset" className={classes}>
      <div className={legendClasses}>
        {legend && <legend>{legend}</legend>}
        {hintText && <Hint>{hintText}</Hint>}
      </div>
      {children}
    </fieldset>
  );
};

Fieldset.propTypes = {
  children: node,
  className: string,
  legendClassName: string,
  legendSrOnly: bool,
  legend: node,
  hintText: string,
};

Fieldset.defaultProps = {
  className: '',
  legendClassName: '',
  legendSrOnly: false,
  legend: null,
  hintText: '',
};

export default Fieldset;
