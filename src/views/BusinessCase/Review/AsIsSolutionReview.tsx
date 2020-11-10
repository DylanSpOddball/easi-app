import React from 'react';

import EstimatedLifecycleCostReview from 'components/EstimatedLifecycleCost/Review';
import ReviewRow from 'components/ReviewRow';
import {
  DescriptionDefinition,
  DescriptionList,
  DescriptionTerm
} from 'components/shared/DescriptionGroup';
import { BusinessCaseSolution } from 'types/businessCase';

/**
 * Title
 * Summary
 * Pros
 * Cons
 * Estimated Lifecycle
 * Cost Savings
 */

type ReviewProps = {
  solution: BusinessCaseSolution;
};

const AsIsSolutionReview = ({ solution }: ReviewProps) => (
  <div>
    <DescriptionList title="As-Is Solution">
      <ReviewRow>
        <div className="line-height-body-3">
          <DescriptionTerm term="“As is” solution: Title" />
          <DescriptionDefinition definition={solution.title} />
        </div>
      </ReviewRow>
      <ReviewRow>
        <div className="line-height-body-3">
          <DescriptionTerm term="“As is” solution: Summary" />
          <DescriptionDefinition
            className="text-pre-wrap"
            definition={solution.summary}
          />
        </div>
      </ReviewRow>
      <ReviewRow>
        <div className="line-height-body-3">
          <DescriptionTerm term="“As is” solution: Pros" />
          <DescriptionDefinition
            className="text-pre-wrap"
            definition={solution.pros}
          />
        </div>
      </ReviewRow>
      <ReviewRow>
        <div className="line-height-body-3">
          <DescriptionTerm term="“As is” solution: Cons" />
          <DescriptionDefinition
            className="text-pre-wrap"
            definition={solution.cons}
          />
        </div>
      </ReviewRow>
    </DescriptionList>
    <ReviewRow>
      <EstimatedLifecycleCostReview data={solution.estimatedLifecycleCost} />
    </ReviewRow>

    <ReviewRow>
      <div className="line-height-body-3">
        <DescriptionTerm term="What is the cost savings or avoidance associated with this solution?" />
        <DescriptionDefinition
          className="text-pre-wrap"
          definition={solution.costSavings}
        />
      </div>
    </ReviewRow>
  </div>
);

export default AsIsSolutionReview;
