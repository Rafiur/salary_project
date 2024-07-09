package grpchandler

import (
	"context"
	"fmt"
	"salary_project/entity"
	"salary_project/entity/proto"
	"salary_project/service"
)

type SalaryGrpcHandler struct {
	proto.UnimplementedEmployeeToSalaryServer
	SalaryService *service.SalaryService
}

func (s *SalaryGrpcHandler) CreateSalary(ctx context.Context, in *proto.CreateSalaryRequest) (*proto.CreateSalaryResponse, error) {

	fmt.Println(in)
	//fmt.Println("DDDDDDDDDDDDDDDDDDDDDDDDD")
	payload := entity.CreateEmployeeSalary{
		Salary_Amount: int(in.SalaryAmount),
		Joining_Date:  in.JoiningDate,
		Project:       in.Project,
		Employee_Id:   in.EmployeeId,
	}
	res, err := s.SalaryService.AddSalaryService(ctx, payload)
	if err != nil {
		fmt.Println(err)
		return &proto.CreateSalaryResponse{}, err
	}

	fmt.Println(res)

	resp := &proto.CreateSalaryResponse{
		Status:  true,
		Message: "Salary Created successfully",
	}

	return resp, nil
}

func (s *SalaryGrpcHandler) UpdateSalary(ctx context.Context, in *proto.UpdateSalaryRequest) (*proto.UpdateSalaryResponse, error) {

	//fmt.Println("printing in:",in)

	payload := entity.CreateEmployeeSalary{
		Salary_Amount: int(in.SalaryAmount),
		Joining_Date:  in.JoiningDate,
		Project:       in.Project,
		Employee_Id:   in.EmployeeId,
	}
	//fmt.Println("Printing payload: ",payload)

	res, err := s.SalaryService.UpdateSalaryByIdService(ctx, payload)
	if err != nil {
		fmt.Println(err)
		return &proto.UpdateSalaryResponse{}, err
	}

	fmt.Println(res)

	resp := &proto.UpdateSalaryResponse{
		EmployeeId:   res.Employee_Id,
		SalaryAmount: int32(res.Salary_Amount),
		Project:      res.Project,
		JoiningDate:  res.Joining_Date,
	}

	return resp, nil
}
