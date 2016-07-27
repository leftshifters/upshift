//
//  FullScreenExample.m
//  FSCalendar
//
//  Created by Wenchao Ding on 9/16/15.
//  Copyright (c) 2015 wenchaoios. All rights reserved.
//

#import "FullScreenExampleViewController.h"

@interface FullScreenExampleViewController()

@property (strong, nonatomic) NSCalendar *lunarCalendar;
@property (strong, nonatomic) NSArray *lunarChars;

@property (strong, nonatomic) NSDate *minimumDate;
@property (strong, nonatomic) NSDate *maximumDate;

@end

@implementation FullScreenExampleViewController

- (instancetype)init
{
    self = [super init];
    if (self) {
        self.title = @"FSCalendar";
        _lunarCalendar = [[NSCalendar alloc] initWithCalendarIdentifier:NSCalendarIdentifierChinese];
        _lunarCalendar.locale = [NSLocale localeWithLocaleIdentifier:@"zh-CN"];
        _lunarChars = @[@"初一",@"初二",@"初三",@"初四",@"初五",@"初六",@"初七",@"初八",@"初九",@"初十",@"十一",@"十二",@"十三",@"十四",@"十五",@"十六",@"十七",@"十八",@"十九",@"二十",@"二一",@"二二",@"二三",@"二四",@"二五",@"二六",@"二七",@"二八",@"二九",@"三十"];
    }
    return self;
}

- (void)loadView
{
    UIView *view = [[UIView alloc] initWithFrame:[UIScreen mainScreen].bounds];
    view.backgroundColor = [UIColor colorWithRed:0.95 green:0.95 blue:0.95 alpha:1.0];
    self.view = view;
    
    FSCalendar *calendar = [[FSCalendar alloc] initWithFrame:CGRectMake(0, 64, self.view.bounds.size.width, self.view.bounds.size.height-64)];
    calendar.backgroundColor = [UIColor whiteColor];
    calendar.dataSource = self;
    calendar.delegate = self;
    calendar.pagingEnabled = NO; // important
    calendar.allowsMultipleSelection = YES;
    calendar.firstWeekday = 2;
    calendar.placeholderType = FSCalendarPlaceholderTypeFillHeadTail;
    calendar.appearance.caseOptions = FSCalendarCaseOptionsWeekdayUsesSingleUpperCase|FSCalendarCaseOptionsHeaderUsesUpperCase;
    calendar.placeholderType = FSCalendarPlaceholderTypeFillHeadTail;
    [self.view addSubview:calendar];
    self.calendar = calendar;
    
    UIBarButtonItem *todayItem = [[UIBarButtonItem alloc] initWithTitle:@"Today" style:UIBarButtonItemStyleBordered target:self action:@selector(todayItemClicked:)];
    
    UIBarButtonItem *lunarItem = [[UIBarButtonItem alloc] initWithTitle:@"Lunar" style:UIBarButtonItemStyleBordered target:self action:@selector(lunarItemClicked:)];
    [lunarItem setTitleTextAttributes:@{NSForegroundColorAttributeName:[UIColor magentaColor]} forState:UIControlStateNormal];
    self.navigationItem.rightBarButtonItems = @[lunarItem, todayItem];
}

- (void)viewDidLoad
{
    [super viewDidLoad];
    self.minimumDate = [self.calendar dateWithYear:2016 month:2 day:1];
    self.maximumDate = [self.calendar dateWithYear:2018 month:4 day:1];
    /*
    dispatch_after(dispatch_time(DISPATCH_TIME_NOW, (int64_t)(3 * NSEC_PER_SEC)), dispatch_get_main_queue(), ^{
        self.minimumDate = [self.calendar dateWithYear:2015 month:2 day:1];
        self.maximumDate = [self.calendar dateWithYear:2015 month:6 day:1];
        [self.calendar reloadData];
    });
    */
}

- (void)dealloc
{
    NSLog(@"%s",__FUNCTION__);
}

- (void)todayItemClicked:(id)sender
{
    [_calendar setCurrentPage:[NSDate date] animated:YES];
}

- (void)lunarItemClicked:(UIBarButtonItem *)item
{
    _showsLunar = !_showsLunar;
    [_calendar reloadData];
}

- (void)viewWillLayoutSubviews
{
    [super viewWillLayoutSubviews];
    _calendar.frame = CGRectMake(0, CGRectGetMaxY(self.navigationController.navigationBar.frame), self.view.bounds.size.width, self.view.bounds.size.height-CGRectGetMaxY(self.navigationController.navigationBar.frame));
}

- (NSDate *)minimumDateForCalendar:(FSCalendar *)calendar
{
    return self.minimumDate;
}

- (NSDate *)maximumDateForCalendar:(FSCalendar *)calendar
{
    return self.maximumDate;
}

- (NSString *)calendar:(FSCalendar *)calendar subtitleForDate:(NSDate *)date
{
    if (_showsLunar) {
        NSInteger day = [_lunarCalendar components:NSCalendarUnitDay fromDate:date].day;
        return _lunarChars[day-1];
    }
    return nil;
}

- (void)calendar:(FSCalendar *)calendar didSelectDate:(NSDate *)date
{
    NSLog(@"did select %@",[calendar stringFromDate:date format:@"yyyy/MM/dd"]);
}

- (void)calendarCurrentPageDidChange:(FSCalendar *)calendar
{
    NSLog(@"did change page %@",[calendar stringFromDate:calendar.currentPage format:@"yyyy-MM"]);
}

@end
