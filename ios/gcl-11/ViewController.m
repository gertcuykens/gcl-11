//
//  ViewController.m
//  gcl-11
//
//  Created by Gert Cuykens on 30/05/14.
//  Copyright (c) 2014 Gert Cuykens. All rights reserved.
//

#import "ViewController.h"
#import "GTLServiceService.h"
#import <GTLTouchStaticLib/GTMOAuth2Authentication.h>
#import <FacebookSDK/FacebookSDK.h>

typedef void(^function)();

static NSString *loadingText = @"Loading...";

@interface ViewController ()

@property (strong, atomic) GTLServiceService *service;
@property (strong, nonatomic) IBOutlet UITextField *txtMessage;
@property (strong, nonatomic) IBOutlet UITextField *txtGraph;
@property (strong, nonatomic) IBOutlet UITextView *txtOutput;
@property (strong, nonatomic) FBRequestConnection *session;
@property (strong, nonatomic) UITapGestureRecognizer *tap;

- (IBAction)btnSubmitHandler:(id)sender;

- (IBAction)btnPublishHandler:(id)sender;

- (BOOL)textViewShouldBeginEditing:(UITextView *)textView;

- (void)login:(function)handler;

- (void)rpcList;

- (void)rpcSubmit;

- (void)sendRequests;

- (void)requestCompleted:(FBRequestConnection *)connection
                 forFbID:(NSString *)fbID
                  result:(id)result
                   error:(NSError *)error;

+ (void)alert:(NSString *)message;

@end

@implementation ViewController

@synthesize service;

- (void)viewDidLoad{
    [super viewDidLoad];
    [self rpcList];
    
    self.tap = [[UITapGestureRecognizer alloc] initWithTarget:self action:@selector(dismissKeyboard)];
    [self.view addGestureRecognizer:self.tap];
    //_txtOutput.delegate = self;
}

-(void)dismissKeyboard {
    //[self resignFirstResponder];
    //[self.view removeGestureRecognizer:self.tap];
    [self.view endEditing:YES];
}

- (BOOL)textViewShouldBeginEditing:(UITextView *)textView {
    NSLog(@"close keyboard");
    //[self.view endEditing:YES];
    return YES;
}

- (void)didReceiveMemoryWarning{[super didReceiveMemoryWarning];}

- (void)rpcSubmit {
    if (!self.service) {service = [[GTLServiceService alloc] init];}
    
    FBAccessTokenData *d=FBSession.activeSession.accessTokenData;
    NSString *token = d.accessToken;
    NSLog(@"token: %@",token);
    
    GTMOAuth2Authentication *auth=[[GTMOAuth2Authentication alloc] init];
    auth.accessToken = token;
    [service setAuthorizer:auth];
    
    GTLServiceMessage *message = [[GTLServiceMessage alloc] init];
    message.message=self.txtMessage.text;
    NSArray *list = [NSArray arrayWithObject:message];
    
    GTLServiceEntity *entity=[[GTLServiceEntity alloc] init];
    entity.list=list;
    
    GTLQueryService *query = [GTLQueryService queryForDatastoreSubmitWithObject:entity];
    [service executeQuery:query completionHandler:^(GTLServiceTicket *ticket, GTLServiceEntity *object, NSError *error) {
        if (error) NSLog(@"error: %@",error);
        [self rpcList];
    }];
    
    self.service=service;
}

- (void)rpcList {
    
    self.txtOutput.text = @"";
    
    if (!self.service) {service = [[GTLServiceService alloc] init];}
    
    FBAccessTokenData *d=FBSession.activeSession.accessTokenData;
    NSString *token = d.accessToken;
    NSLog(@"token: %@",token);
    
    GTMOAuth2Authentication *auth=[[GTMOAuth2Authentication alloc] init];
    auth.accessToken = token;
    [service setAuthorizer:auth];
    
    GTLServiceEntity *entity=[[GTLServiceEntity alloc] init];

    GTLQueryService *query = [GTLQueryService queryForDatastoreListWithObject:entity];
    [service executeQuery:query
        completionHandler:
        ^(GTLServiceTicket *ticket, GTLServiceEntity *object, NSError *error) {
            if (error) NSLog(@"error: %@",error);
            for (GTLServiceMessage *m in object.list) {
                NSLog(@"object: %@",m.message);
                self.txtOutput.text = [NSString stringWithFormat:@"%@%@\r\n",self.txtOutput.text, m.message];
            }
        }
     ];
    
    self.service=service;
}

- (void)login:(function)handler {
    [FBSession openActiveSessionWithReadPermissions:nil
                                       allowLoginUI:YES
                                  completionHandler:
                                  ^(FBSession *session, FBSessionState status, NSError *error) {
                                      if (error) { [ViewController alert:error.localizedDescription];}
                                      else if (FB_ISSESSIONOPENWITHSTATE(status)) {handler();}
                                  }];
}

- (void)btnSubmitHandler:(id)sender {
    if (FBSession.activeSession.isOpen) {}
    else {[self rpcSubmit];}
}

- (void)btnPublishHandler:(id)sender {
    function f = ^(){[self sendRequests];};
    if (FBSession.activeSession.isOpen) {[self sendRequests];}
    else {[self login:f];}
}

- (void)sendRequests {

    NSArray *fbids = [self.txtGraph.text componentsSeparatedByString:@","];
    
    if ([self.txtGraph.text length] == 0) {[ViewController alert:@"Object ID is required"];return;}
    
    FBRequestConnection *connection = [[FBRequestConnection alloc] init];
    
    for (NSString *fbid in fbids) {
        FBRequestHandler handler =
        ^(FBRequestConnection *connection, id result, NSError *error) {
            [self requestCompleted:connection forFbID:fbid result:result error:error];
        };
        
        NSMutableDictionary *params = [NSMutableDictionary dictionaryWithObjectsAndKeys:
                                       //@"http:", @"link",
                                       //@"my profile", @"name",
                                       //YourUIImage, @"source",
                                       //@"caption description",@"message",
                                       self.txtMessage.text, @"message",
                                       nil];
        
        FBRequest *request1 = [[FBRequest alloc] initWithSession:FBSession.activeSession
                                                       graphPath:fbid
                                                      parameters:params
                                                      HTTPMethod:@"POST"];
        
        [connection addRequest:request1
             completionHandler:handler];
    }
    
    [self.session cancel];
    self.session = connection;
    [connection start];
}

- (void)requestCompleted:(FBRequestConnection *)connection
                 forFbID:fbID
                  result:(id)result
                   error:(NSError *)error {
    if (error) NSLog(@"error: %@", error.localizedDescription);
    NSLog(@"object: %@",[fbID stringByTrimmingCharactersInSet:[NSCharacterSet whitespaceAndNewlineCharacterSet]]);
    NSDictionary *dictionary = (NSDictionary *)result;
    NSLog(@"object: %@",(NSString *)[dictionary objectForKey:@"name"]);
    if (self.session && connection != self.session) {return;}
    self.session = nil;
}

+ (void)alert:message {
    UIAlertView *alert = [[UIAlertView alloc] initWithTitle:@"Error"
                                                    message:message
                                                   delegate:nil
                                          cancelButtonTitle:@"OK"
                                          otherButtonTitles:nil];
    [alert show];
}

@end

//self.txtOutput.text = loadingText;
//if ([self.txtGraph isFirstResponder]) {[self.txtGraph resignFirstResponder];}
//if ([self.txtOutput.text isEqualToString:loadingText]) {self.txtOutput.text = @"";}

//service.retryEnabled = YES;
//service.authorizer = auth;
//[GTMHTTPFetcher setLoggingEnabled:YES];

//FBRequest *request2 = [FBRequest requestForPostStatusUpdate:self.txtMessage.text];

/*
 if (![facebook isSessionValid]) {
 NSArray *permissions = [NSArray arrayWithObjects:@"user_photos",@"user_videos",@"publish_stream",@"offline_access",@"user_checkins",@"friends_checkins",@"email",@"user_location" ,nil];
 [facebook authorize:permissions];
 }
 */
