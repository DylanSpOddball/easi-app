import React from 'react';
import { useTranslation } from 'react-i18next';
import { Link } from '@trussworks/react-uswds';

import BreadcrumbNav from 'components/BreadcrumbNav';
import PageHeading from 'components/PageHeading';

import './index.scss';

const TestingTemplates = () => {
  const { t } = useTranslation('accessibility');
  const vpatConformanceLevels: { name: string; description: string }[] = t(
    'testingTemplates.vpatSection.subSection.item1.levels',
    {
      returnObjects: true
    }
  );
  const testPlanList: string[] = t(
    'testingTemplates.testPlanSection.itemsToProvide',
    {
      returnObjects: true
    }
  );

  const remediationPlanList: string[] = t(
    'testingTemplates.remediationPlanSection.itemsToProvide',
    {
      returnObjects: true
    }
  );

  const tableOfContents = (
    <div className=".accessibility-testing-templates__tableOfContents">
      <p>Page contents</p>
      <ul className="accessibility-testing-templates__table-of-contents">
        <li>
          <Link
            href={`#testingTemplates_${t(
              'testingTemplates.vpatSection.heading'
            )}`}
          >
            {t('testingTemplates.vpatSection.heading')}
          </Link>
        </li>
        <li>
          <Link
            href={`#testingTemplates_${t(
              'testingTemplates.testPlanSection.heading'
            )}`}
          >
            {t('testingTemplates.testPlanSection.heading')}
          </Link>
        </li>
        <li>
          <Link
            href={`#testingTemplates_${t(
              'testingTemplates.remediationPlanSection.heading'
            )}`}
          >
            {t('testingTemplates.remediationPlanSection.heading')}
          </Link>
        </li>
      </ul>
    </div>
  );

  return (
    <div className="grid-container accessibility-testing-templates">
      <BreadcrumbNav className="margin-y-2">
        <li>
          <Link href="/">Home</Link>
          <i className="fa fa-angle-right margin-x-05" aria-hidden />
        </li>
        <li>Templates for 508 testing</li>
      </BreadcrumbNav>
      <div className="grid-row grid-gap-lg">
        <div className="grid-col-9 line-height-body-4">
          <div className="tablet:grid-col-10">
            <PageHeading>{t('testingTemplates.heading')}</PageHeading>
            {tableOfContents}
            <h2
              id={`testingTemplates_${t(
                'testingTemplates.vpatSection.heading'
              )}`}
            >
              {t('testingTemplates.vpatSection.heading')}
            </h2>
            <p>{t('testingTemplates.vpatSection.description')}</p>
            <h3>{t('testingTemplates.vpatSection.subSection.heading')}</h3>
            <ul className="accessibility-testing-templates__vpat-list">
              <li>{t('testingTemplates.vpatSection.subSection.item1.text')}</li>
              {vpatConformanceLevels.map(level => (
                <p>
                  <span className="text-bold">{level.name}</span>{' '}
                  {level.description}
                </p>
              ))}
              <li>{t('testingTemplates.vpatSection.subSection.item2.text')}</li>
            </ul>
            <div className="accessibility-testing-templates__downloadBox">
              <h3 className="margin-top-0">
                {t(
                  'testingTemplates.vpatSection.subSection.downloadVPAT.heading'
                )}
              </h3>
              <p>
                <Link
                  href="https://www.itic.org/policy/accessibility/vpat"
                  target="_blank"
                >
                  {t(
                    'testingTemplates.vpatSection.subSection.downloadVPAT.line1.linkText'
                  )}
                </Link>
                {` `}
                {t(
                  'testingTemplates.vpatSection.subSection.downloadVPAT.line1.otherText'
                )}
              </p>
              <p>
                <i className="fa fa-exclamation-circle fa-2x" />
                {` `}
                {t(
                  'testingTemplates.vpatSection.subSection.downloadVPAT.line2.text'
                )}
              </p>
              <p>
                <Link
                  href="https://www.youtube.com/watch?v=kAkSV9xiJ1A"
                  target="_blank"
                >
                  {t(
                    'testingTemplates.vpatSection.subSection.downloadVPAT.line3.linkText'
                  )}
                </Link>
                {` `}
                {t(
                  'testingTemplates.vpatSection.subSection.downloadVPAT.line3.otherText'
                )}
              </p>
            </div>
            <div>
              <h2
                id={`testingTemplates_${t(
                  'testingTemplates.testPlanSection.heading'
                )}`}
              >
                {t('testingTemplates.testPlanSection.heading')}
              </h2>
              <p>{t('testingTemplates.testPlanSection.description')}</p>
              <ul className="accessibility-testing-templates__test-plan-list">
                {testPlanList.map(item => (
                  <li>{item}</li>
                ))}
              </ul>
            </div>
            <div>
              <h2
                id={`testingTemplates_${t(
                  'testingTemplates.remediationPlanSection.heading'
                )}`}
              >
                {t('testingTemplates.remediationPlanSection.heading')}
              </h2>
              <p>{t('testingTemplates.remediationPlanSection.description')}</p>
              <ul className="accessibility-testing-templates__remediation-plan-list">
                {remediationPlanList.map(item => (
                  <li>{item}</li>
                ))}
              </ul>
            </div>
          </div>
        </div>
        <div className="grid-col-3 accessibility-testing-templates__side-col">
          <div>
            <h4 className="text-bold">
              Need help? Contact the Section 508 team
            </h4>
            <Link href="mailto:CMS_Section508@cms.hhs.gov">
              CMS_Section508@cms.hhs.gov
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
};

export default TestingTemplates;
