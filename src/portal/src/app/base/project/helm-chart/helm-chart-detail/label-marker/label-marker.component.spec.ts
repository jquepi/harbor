import { ComponentFixture, TestBed } from '@angular/core/testing';
import { TranslateModule, TranslateService } from '@ngx-translate/core';
import { LabelMarkerComponent } from './label-marker.component';
import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { ClarityModule } from '@clr/angular';
import { FormsModule } from '@angular/forms';
import { of } from 'rxjs';
import { LabelService } from '../../../../../shared/services';
import { ErrorHandler } from '../../../../../shared/units/error-handler';

describe('LabelMarkerComponent', () => {
    const mockErrorHandler = null;

    const mockLabelService = {
        getChartVersionLabels: () => {
            return of({
                name: '111',
                description: 'string',
                color: 'string',
                scope: 'string',
                project_id: 1,
            });
        },
        markChartLabel: () => {},
        unmarkChartLabel: () => {},
    };
    let component: LabelMarkerComponent;
    let fixture: ComponentFixture<LabelMarkerComponent>;

    beforeEach(() => {
        TestBed.configureTestingModule({
            schemas: [CUSTOM_ELEMENTS_SCHEMA],
            imports: [ClarityModule, TranslateModule.forRoot(), FormsModule],
            declarations: [LabelMarkerComponent],
            providers: [
                TranslateService,
                { provide: LabelService, useValue: mockLabelService },
                { provide: ErrorHandler, useValue: mockErrorHandler },
            ],
        }).compileComponents();
    });

    beforeEach(() => {
        fixture = TestBed.createComponent(LabelMarkerComponent);
        component = fixture.componentInstance;
        fixture.detectChanges();
    });

    it('should create', () => {
        expect(component).toBeTruthy();
    });
});
