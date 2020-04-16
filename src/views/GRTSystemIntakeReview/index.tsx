import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RouteComponentProps } from 'react-router-dom';
import { SystemIntakeReview } from 'components/SystemIntakeReview';
import { getSystemIntake } from '../../actions/systemIntakeActions';
import { AppState } from '../../reducers/rootReducer';

export type SystemIDRouterProps = {
  systemId: string;
};

export type GRTSystemIntakeReviewProps = RouteComponentProps<
  SystemIDRouterProps
> & {
  auth: any;
};

export const GRTSystemIntakeReview = ({
  match
}: GRTSystemIntakeReviewProps) => {
  const dispatch = useDispatch();

  useEffect(() => {
    const fetchSystemIntake = async (): Promise<void> => {
      dispatch(getSystemIntake(await match.params.systemId));
    };
    fetchSystemIntake();
  }, [dispatch, match.params.systemId]);
  const heading = 'CMS Sytem Request';
  const systemIntake = useSelector(
    (state: AppState) => state.systemIntake.systemIntake
  );
  return (
    <div>
      <SystemIntakeReview
        heading={heading}
        systemIntake={systemIntake}
        id={match.params.systemId}
      />
      <hr className="system-intake__hr" />
      <h2 className="font-heading-xl">What to do after reviewing?</h2>
      <p>Email the requester and:</p>
      <ul className="usa-list">
        <li>Ask them to fill in the business case and submit it</li>
        <li>Tell them what to expect after they submit the business case</li>
        <li>And how to get in touch if they have any questions.</li>
      </ul>
    </div>
  );
};

export default GRTSystemIntakeReview;
