import { of as observableOf } from 'rxjs';
import { DebugElement } from '@angular/core';
import { TestBed, ComponentFixture } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { StoreModule, Store } from '@ngrx/store';
import { MockComponent } from 'ng2-mock-component';

import { using } from 'app/testing/spec-helpers';
import { FeatureFlagsService } from 'app/services/feature-flags/feature-flags.service';
import { SettingsLandingComponent } from 'app/pages/settings-landing/settings-landing.component';
import { NgrxStateAtom, ngrxReducers, defaultInitialState, runtimeChecks } from 'app/ngrx.reducers';
import { checkFirstPerm } from 'app/testing/spec-helpers';
import { SettingsSidebarComponent } from './settings-sidebar.component';

describe('SettingsSidebarComponent', () => {
  let store: Store<NgrxStateAtom>;
  let fixture: ComponentFixture<SettingsSidebarComponent>;
  let component: SettingsSidebarComponent;
  let settingsLandingComponent: SettingsLandingComponent;
  let element: DebugElement;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [
        RouterTestingModule,
        StoreModule.forRoot(ngrxReducers, { initialState: defaultInitialState, runtimeChecks })
      ],
      declarations: [
        SettingsSidebarComponent,
        SettingsLandingComponent,
        // Note: if we want to match contents, we cannot swallow them: thus we
        // need to provide a template here. <ng-content> mocks these as doing
        // nothing but a "pass-through" of what the components wrap.
        MockComponent({ selector: 'app-landing',
                        inputs: ['routePerms'] }),
        MockComponent({ selector: 'app-authorized',
                        inputs: ['allOf', 'anyOf'],
                        template: '<ng-content></ng-content>' }),
        MockComponent({ selector: 'chef-sidebar',
                        template: '<ng-content></ng-content>' }),
        MockComponent({ selector: 'chef-sidebar-entry',
                        inputs: ['route', 'icon', 'exact'],
                        template: '<ng-content></ng-content>' })
      ],
      providers: [
        FeatureFlagsService
      ]
    });
    store = TestBed.get(Store);
    spyOn(store, 'dispatch').and.callThrough();
    fixture = TestBed.createComponent(SettingsSidebarComponent);
    component = fixture.componentInstance;
    component.featureFlagOn = true;
    settingsLandingComponent =
      TestBed.createComponent(SettingsLandingComponent).componentInstance;
    element = fixture.debugElement;
  });

  it('should be created', () => {
    fixture.detectChanges();

    expect(component).toBeTruthy();
  });

  describe('IAM v2', () => {
    it('shows all links consistent with settings-landing', () => {
      component.isIAMv2$ = observableOf(true);
      fixture.detectChanges();
      const links = element.nativeElement
        .querySelectorAll('div.nav-items chef-sidebar-entry');
      expect(links.length).toBe(11);
    });

    it('has route order consistent with settings-landing', () => {
      component.isIAMv2$ = observableOf(true);
      fixture.detectChanges();
      const links = element.nativeElement
        .querySelectorAll('div.nav-items chef-sidebar-entry');
      for (let i = 0; i < links.length; i++) {
        expect(links[i].getAttribute('route'))
          .toBe(settingsLandingComponent.routeList[i].route);
      }
    });

    it('has paths consistent with settings-landing', () => {
      component.isIAMv2$ = observableOf(true);
      fixture.detectChanges();
      const elements = Array
        .from<HTMLElement>(
          element.nativeElement.querySelectorAll('div.nav-items app-authorized'))
        .filter(elem =>
          elem.firstElementChild && elem.firstElementChild.tagName === 'CHEF-SIDEBAR-ENTRY');

      for (let i = 0; i < elements.length; i++) {
        checkFirstPerm(
          'anyOf',
          elements[i].getAttribute('ng-reflect-any-of'),
          settingsLandingComponent.routeList[i].anyOfCheck);
        checkFirstPerm(
          'allOf',
          elements[i].getAttribute('ng-reflect-all-of'),
          settingsLandingComponent.routeList[i].allOfCheck);
      }
    });
  });

  describe('IAM v1', () => {
    it('shows 8 links', () => {
      component.isIAMv2$ = observableOf(false);
      fixture.detectChanges();
      const links = element.nativeElement
        .querySelectorAll('div.nav-items chef-sidebar-entry');
      expect(links.length).toBe(8);
    });

    using([
      ['Notifications', '/settings/notifications', 0],
      ['Data Feeds', '/settings/data-feed', 1],
      ['Node Integrations', '/settings/node-integrations', 2],
      ['Node Credentials', '/settings/node-credentials', 3],
      ['Node Lifecycle', '/settings/node-lifecycle', 4],
      ['Users', '/settings/users', 5],
      ['Teams', '/settings/teams', 6],
      ['API Tokens', '/settings/tokens', 7]
    ], (label: string, path: string, position: number) => {
      it(`displays the ${label} navigation link`, () => {
        component.isIAMv2$ = observableOf(false);
        fixture.detectChanges();
        const links = element.nativeElement
          .querySelectorAll('div.nav-items chef-sidebar-entry');
        expect(links[position].innerText).toBe(label);
        expect(links[position].getAttribute('route')).toBe(path);
      });
    });
  });
});
