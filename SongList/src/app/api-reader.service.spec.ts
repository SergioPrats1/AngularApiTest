import { TestBed } from '@angular/core/testing';

import { ApiReaderService } from './api-reader.service';

describe('ApiReaderService', () => {
  let service: ApiReaderService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ApiReaderService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
