import { Component, OnInit } from '@angular/core';
import { InspectorSearchGQL, InspectorSearchQuery, SearchFilter, SortOrder } from "../../../generated/graphql";
import { debounceTime, map, Observable, startWith, tap } from "rxjs";
import { FormControl, FormGroup } from "@angular/forms";

@Component({
  selector: 'app-inspector',
  templateUrl: './inspector.component.html',
  styleUrls: ['./inspector.component.scss']
})
export class InspectorComponent implements OnInit {
  public loading: boolean = true;
  public empty: boolean = false;
  public result$: Observable<InspectorSearchQuery> | undefined;

  public form: FormGroup = new FormGroup({});
  public pathCtrl: FormControl = new FormControl();
  public pageCtrl: FormControl = new FormControl(0);
  public lowerRatingCtrl: FormControl = new FormControl();
  public upperRatingCtrl: FormControl = new FormControl();
  public sortOrderCtrl: FormControl = new FormControl();

  public sortOrderOptions: { name: string; id: SortOrder }[] = [];

  private pageSize: number = 11;

  constructor(private query: InspectorSearchGQL) {
    this.form.addControl('pathCtrl', this.pathCtrl);
    this.form.addControl('pageCtrl', this.pageCtrl);
    this.form.addControl('lowerRatingCtrl', this.lowerRatingCtrl);
    this.form.addControl('upperRatingCtrl', this.upperRatingCtrl);
    this.form.addControl('sortOrderCtrl', this.sortOrderCtrl);
  }

  ngOnInit(): void {
    const savedForm = localStorage.getItem('form');
    if (savedForm && savedForm.length > 3) {
      this.form.setValue(JSON.parse(savedForm))
    }

    this.sortOrderOptions = Object.keys(SortOrder)
      .filter(v => isNaN(Number(v)))
      .map(name => ({id: SortOrder[name as keyof typeof SortOrder], name}));

    this.form.valueChanges.pipe(
      startWith(this.createFilter()),
      debounceTime(500)).subscribe(() => this.update());

    this.query.watch({input: this.createFilter()}).refetch();
  }

  private update(): void {
    const filter = this.createFilter();

    let watch = this.query.watch({input: filter});
    this.result$ = watch.valueChanges.pipe(
      tap(res => this.loading = res.loading),
      map(res => res.data),
      tap(data => this.empty = data.Pictures?.length === 0),
    );
  }

  private createFilter(): SearchFilter {
    let filter: SearchFilter = {
      lowerRating: 0,
      take: this.pageSize,
      skip: 0,
      sortOrder: this.sortOrderCtrl.value,
    };

    if (this.lowerRatingCtrl.value != 0) {
      filter.lowerRating = this.lowerRatingCtrl.value;
    }

    if (this.upperRatingCtrl.value != 0) {
      filter.upperRating = this.upperRatingCtrl.value;
    }

    if (this.pageCtrl.value != 0) {
      filter.take = this.pageSize;
      filter.skip = this.pageSize * this.pageCtrl.value;
    }

    if (this.pathCtrl.value != '') {
      filter.pathContains = this.pathCtrl.value;
    }

    console.log('created filter', filter);
    localStorage.setItem('form', JSON.stringify(this.form.value));
    return filter
  }

  public previousPageClicked(): void {
    if (this.pageCtrl.value == 0) {
      return;
    }

    this.pageCtrl.setValue(this.pageCtrl.value - 1);
  }

  public nextPageClicked(): void {
    this.pageCtrl.setValue(this.pageCtrl.value + 1)
  }

  public resetPage(): void {
    this.pageCtrl.setValue(0);
  }

  public resetFormFilter(): void {
    this.form.reset();
    this.pageCtrl.setValue(0);
  }
}
