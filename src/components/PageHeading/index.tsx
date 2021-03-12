import React, { useEffect, useRef } from 'react';
import classnames from 'classnames';

import './index.scss';

type PageHeadingProps = {
  children: React.ReactNode;
  className?: string;
} & JSX.IntrinsicElements['h1'];

/**
 * This is h1 that belongs on every page.
 * Design wants to standardize the margins around h1 that appear at the top of the page.
 * This gives the h1 element more room to breathe.
 */
const PageHeading = ({ children, className, ...props }: PageHeadingProps) => {
  const headingRef = useRef<HTMLHeadingElement>(null);
  const classes = classnames(
    'easi-h1',
    'margin-top-6',
    'margin-bottom-5',
    className
  );

  useEffect(() => {
    headingRef.current?.focus();
  }, []);

  return (
    <h1 className={classes} tabIndex={-1} ref={headingRef} {...props}>
      {children}
    </h1>
  );
};

export default PageHeading;
