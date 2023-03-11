// @generated by protoc-gen-connect-web v0.8.3 with parameter "target=ts"
// @generated from file group.proto (package proto, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { DeleteGroupRequest, DeleteGroupResponse, GetGroupRequest, GetGroupResponse, ListGroupRequest, ListGroupResponse, RegisterGroupRequest, RegisterGroupResponse, UpdateGroupRequest, UpdateGroupResponse } from "./group_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service proto.GroupService
 */
export const GroupService = {
  typeName: "proto.GroupService",
  methods: {
    /**
     * @generated from rpc proto.GroupService.ListGroup
     */
    listGroup: {
      name: "ListGroup",
      I: ListGroupRequest,
      O: ListGroupResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.GroupService.GetGroup
     */
    getGroup: {
      name: "GetGroup",
      I: GetGroupRequest,
      O: GetGroupResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.GroupService.RegisterGroup
     */
    registerGroup: {
      name: "RegisterGroup",
      I: RegisterGroupRequest,
      O: RegisterGroupResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.GroupService.UpdateGroup
     */
    updateGroup: {
      name: "UpdateGroup",
      I: UpdateGroupRequest,
      O: UpdateGroupResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc proto.GroupService.DeleteGroup
     */
    deleteGroup: {
      name: "DeleteGroup",
      I: DeleteGroupRequest,
      O: DeleteGroupResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

